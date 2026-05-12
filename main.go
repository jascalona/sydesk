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

	// INYECCION DE DEPENDENCIAS (CAPA DE DATOS -> SERVICIO -> HANDLER)

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

	// AUDITORIA
	auditRepo := repoAudit.NewAuditRepo(dbConn)
	asRepo := repoAudit.NewAsRepo(dbConn)
	channelRepo := repoAudit.NewChannelRepo(dbConn)
	supRepo := repoAudit.NewSupRepo(dbConn)

	// Grupo negocio
	customerRepo := repoBus.NewCustomerRepo(dbConn)

	// Grupo de Servicios
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

	customerService := servBus.NewCustomerServ(customerRepo)
	auditService := servAudit.NewAuditServ(auditRepo)
	asService := servAudit.NewAsServ(asRepo)
	channelService := servAudit.NewChannelServ(channelRepo)
	supService := servAudit.NewSupService(supRepo)

	// grupo de servicios bloqueados
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

	// grupo de negocio
	auditHandler := handler.NewAuditHandler(auditService)
	asHandler := handler.NewAsHandler(asService)
	customerHandler := handler.NewCustomerHandler(customerService)
	channelHandler := handler.NewChannelHandler(channelService)
	supHandler := handler.NewSupHandler(supService)

	// Servicio de Autenticación (Login/JWT)
	// Obtenemos la llave secreta desde el .env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "clave_por_defecto_solo_para_dev"
	}
	authService := middlerware.NewAuthService(authRepo, jwtSecret)

	// INICIALIZACIÓN DE GIN Y RUTAS
	r := gin.Default()

	// Llamamos al Administrador Central de Rutas
	// Pasamos el router (r), el servicio de auth y los handlers de cada módulo
	router.SetupRouter(
		r,
		authService,
		userHandler,
		rolesHandler,
		productHandler,
		componentHandler,
		subcomponentHandler,
		enviromentHandler,
		statusHandler,
		customerHandler,
		customRoleHandler,
		customPRHandler,
		auditHandler,
		asHandler,
		channelHandler,
		supHandler,
	)

	// EJECUCION DEL SERVIDOR
	port := ":8081"
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
