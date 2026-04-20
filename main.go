package main

import (
	"log"
	"os"
	"sydesk/config"
	"sydesk/internal/database"
	"sydesk/internal/handler"
	"sydesk/internal/handler/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Imports de repositorios y servicios
	repOpg "sydesk/internal/repository/organization"
	servOrg "sydesk/internal/service/organization"
)

func main() {
	// 1. Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Archivo .env no encontrado, usando variables de sistema.")
	}

	// 2. Cargar configuración e inicializar DB
	cfg := config.LoadConfig()
	dbConn := database.InitDB(cfg.DatabaseURL)
	defer dbConn.Close()

	// 3. INYECCIÓN DE DEPENDENCIAS (CAPA DE DATOS -> SERVICIO -> HANDLER)

	// Repositorio base (se comparte entre servicios)
	userRepo := repOpg.NewUserRepo(dbConn)

	// Servicio de Usuarios (CRUD)
	userService := servOrg.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Servicio de Autenticación (Login/JWT)
	// Obtenemos la llave secreta desde el .env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "clave_por_defecto_solo_para_dev"
	}
	authService := servOrg.NewAuthService(userRepo, jwtSecret)

	// 4. INICIALIZACIÓN DE GIN Y RUTAS
	r := gin.Default()

	// Llamamos al Administrador Central de Rutas
	// Pasamos el router (r), el servicio de auth y los handlers de cada módulo
	router.SetupRouter(r, authService, userHandler)

	// 5. EJECUCIÓN DEL SERVIDOR
	port := ":8081"
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
