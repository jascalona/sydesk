package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Librería para leer el .env
	"log"
	"sydesk/config"
	"sydesk/internal/database"
	"sydesk/internal/handler"
	"sydesk/internal/handler/router"

	// imports section services
	repOpg "sydesk/internal/repository/organization"
	servOrg "sydesk/internal/service/organization"
)

func main() {

	// environment variable
	if err := godotenv.Load(); err != nil {
		log.Println("The configuration file was not found.\n")
	}
	// load config
	cfg := config.LoadConfig()
	// init security connection
	dbConn := database.InitDB(cfg.DatabaseURL)
	// security close pool connection
	defer dbConn.Close()

	// init starting gin
	r := gin.Default()

	// injection dependence

	/*** SECTION ORGANIZATION ***/
	userHandler := handler.NewUserHandler(servOrg.NewUserService(repOpg.NewUserRepo(dbConn)))

	router.ConfigUserRouter(r, userHandler)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}

}
