package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var db *sql.DB

func InitDB(databaseURL string) *sql.DB {
	var err error
	db, err = sql.Open("postgres", databaseURL)

	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Printf("Re-attempting  connection")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("The connection cloud not be established")
	}

	fmt.Println("Successfully connected to the database")
	return db
}

// close pool the connections
func CloseDB() {
	if db != nil {
		fmt.Println("Closing the database successfully")
	}
}
