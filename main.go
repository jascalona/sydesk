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
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Archivo .env no encontrado, usando variables de sistema.")
	}

	// Cargar configuración e inicializar DB
	cfg := config.LoadConfig()
	dbConn := database.InitDB(cfg.DatabaseURL)
	defer dbConn.Close()

	// INYECCIÓN DE DEPENDENCIAS (CAPA DE DATOS -> SERVICIO -> HANDLER)

	authRepo := middlerware.NewAuthRepo(dbConn)

	// Grupo de repositorios bloqueados
	userRepo := repOpg.NewUserRepo(dbConn)

	// Grupo componentes
	productRepo := repoCom.NewProductRepo(dbConn)
	componentRepo := repoCom.NewComponentsRepo(dbConn)

	// Grupo de Servicios
	userService := servOrg.NewUserService(userRepo)

	// Grupo de servicios componentes
	productService := servComp.NewProductService(productRepo)
	componentService := servComp.NewComponentService(componentRepo)

	// grupo de servicios bloqueados
	userHandler := handler.NewUserHandler(userService)

	// grupo de componentes
	productHandler := handler.NewProductHandler(productService)
	componentHandler := handler.NewComponentsHandler(componentService)

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
		productHandler,
		componentHandler,
	)

	// EJECUCION DEL SERVIDOR
	port := ":8081"
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
