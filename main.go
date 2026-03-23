package main

import (
	"github.com/joho/godotenv" // Librería para leer el .env
	"log"
	"sydesk/config"
	"sydesk/internal/database"
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

}
