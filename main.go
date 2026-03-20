package main

import (
	"sydesk/config"
	"sydesk/internal/database"
)

func main() {
	// load condfig
	conf := config.LoadConfig()
	dbConn := database.InitDB(conf.DatabaseURL)
	defer dbConn.Close()

	// hay que implementar toda la validacion del main
}
