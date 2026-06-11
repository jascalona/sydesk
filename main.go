package main

import (
	"log"
	"os"
	"sydesk/config"
	"sydesk/internal/database"
	"sydesk/internal/handler"
	"sydesk/internal/handler/router"
	"sydesk/internal/middlerware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Imports de repositorios y servicios
	repOpg "sydesk/internal/repository/organization"
	servOrg "sydesk/internal/service/organization"

	repoCom "sydesk/internal/repository/components"
	servComp "sydesk/internal/service/components"

	repoBus "sydesk/internal/repository/business"
	servBus "sydesk/internal/service/business"

	repoAudit "sydesk/internal/repository/audit"
	servAudit "sydesk/internal/service/audit"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Archivo .env no encontrado, usando variables de sistema.")
	}

	// Cargar configuracion e inicializar DB
	cfg := config.LoadConfig()
	dbConn := database.InitDB(cfg.DatabaseURL)
	defer dbConn.Close()

	// =========================================================================
	// INYECCIÓN DE DEPENDENCIAS (CAPA DE DATOS -> REPOSITORIOS)
	// =========================================================================
	authRepo := middlerware.NewAuthRepo(dbConn)

	// Grupo de repositorios bloqueados
	userRepo := repOpg.NewUserRepo(dbConn)
	rolesRepo := repOpg.NewRoleRepo(dbConn)

	// Grupo componentes
	productRepo := repoCom.NewProductRepo(dbConn)
	componentRepo := repoCom.NewComponentsRepo(dbConn)
	subcomponentRepo := repoCom.NewSubcomponentRepo(dbConn)
	enviromentRepo := repoCom.NewEnviromentRepo(dbConn)
	statusRepo := repoCom.NewStatusRepo(dbConn)
	customRole := repoCom.NewCustomRoleRepo(dbConn)
	customPR := repoCom.NewCustomProductRoleRepo(dbConn)
	tyrequest := repoCom.NewTypeRequestRepo(dbConn)

	// AUDITORIA
	auditRepo := repoAudit.NewAuditRepo(dbConn)
	asRepo := repoAudit.NewAsRepo(dbConn)
	channelRepo := repoAudit.NewChannelRepo(dbConn)
	supRepo := repoAudit.NewSupRepo(dbConn)
	itemsRepo := repoAudit.NewItemRepo(dbConn)

	// Grupo negocio
	customerRepo := repoBus.NewCustomerRepo(dbConn)
	contactRepo := repoBus.NewContactRepo(dbConn)

	// =========================================================================
	// INYECCIÓN DE DEPENDENCIAS (CAPA DE NEGOCIO -> SERVICIOS)
	// =========================================================================

	userService := servOrg.NewUserService(userRepo)
	rolesService := servOrg.NewRoleService(rolesRepo)

	// Grupo de servicios componentes
	productService := servComp.NewProductService(productRepo)
	componentService := servComp.NewComponentService(componentRepo)
	subcomponentService := servComp.NewSubcomponentService(subcomponentRepo)
	enviromentService := servComp.NewEnviromentService(enviromentRepo)
	statusService := servComp.NewStatusService(statusRepo)
	customRoleService := servComp.NewCustomRoleService(customRole)
	customPRService := servComp.NewCustomProductRoleService(customPR)
	tyrequestService := servComp.NewTypeRequestService(tyrequest)

	customerService := servBus.NewCustomerServ(customerRepo)
	contactService := servBus.NewContactService(contactRepo)

	auditService := servAudit.NewAuditServ(auditRepo)
	asService := servAudit.NewAsServ(asRepo)
	channelService := servAudit.NewChannelServ(channelRepo)
	supService := servAudit.NewSupService(supRepo)
	itemsService := servAudit.NewItemService(itemsRepo)

	// Servicio de Autenticación (Login/JWT)
	// Obtenemos la llave secreta desde el .env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "clave_por_defecto_solo_para_dev"
	}
	authService := middlerware.NewAuthService(authRepo, jwtSecret)

	// =========================================================================
	// INYECCIÓN DE DEPENDENCIAS (CAPA DE PRESENTACIÓN -> HANDLERS)
	// =========================================================================

	userHandler := handler.NewUserHandler(userService)
	rolesHandler := handler.NewRolesHandler(rolesService)

	// grupo de componentes
	productHandler := handler.NewProductHandler(productService)
	componentHandler := handler.NewComponentsHandler(componentService)
	subcomponentHandler := handler.NewSubcomponentHandler(subcomponentService)
	enviromentHandler := handler.NewEnviromentHandler(enviromentService)
	statusHandler := handler.NewStatusHandler(statusService)
	customRoleHandler := handler.NewCustomRoleHandler(customRoleService)
	customPRHandler := handler.NewCustomerProductRoleHandler(customPRService)
	tyrequestHandler := handler.NewTypeRequestHandler(tyrequestService)

	customerHandler := handler.NewCustomerHandler(customerService)
	contactHandler := handler.NewContactHandler(contactService)

	// grupo de negocio
	auditHandler := handler.NewAuditHandler(auditService)
	asHandler := handler.NewAsHandler(asService)
	channelHandler := handler.NewChannelHandler(channelService)
	supHandler := handler.NewSupHandler(supService)
	itemsHandler := handler.NewItemHandler(itemsService)

	// =========================================================================
	// INSTANCIACIÓN DE ENRUTADORES MODULARES
	// =========================================================================

	compRouter := router.NewRouterComponent(
		componentHandler,
		customRoleHandler,
		customPRHandler,
		enviromentHandler,
		productHandler,
		statusHandler,
		subcomponentHandler,
		tyrequestHandler,
	)

	orgRouter := router.NewRouterOrganization(userHandler, rolesHandler)

	businessRouter := router.NewRouterBusiness(customerHandler, contactHandler)

	auditRouter := router.NewRouterAudit(
		auditHandler,
		asHandler,
		channelHandler,
		itemsHandler,
		supHandler,
	)

	// Agrupamos todos los módulos en la estructura principal del Router
	apiRouters := router.MainRouters{
		ComponetRouter:     compRouter,
		BusinessRouter:     businessRouter,
		OrganizationRouter: orgRouter,
		AuditRouter:        auditRouter,
	}

	// =========================================================================
	// INICIALIZACIÓN DE GIN Y RUTAS CENTRALES
	// =========================================================================
	r := gin.Default()

	// Arrancamos el administrador central con los enrutadores empaquetados
	router.SetupRouter(r, authService, apiRouters)

	// se invoca a la funcion para los registros de usuarios para los casos de registros
	router.RegisterUser(r, userHandler)

	// EJECUCION DEL SERVIDOR
	port := ":8081"
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
