package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var db *sql.DB

func InitDB(databaseURL string) *sql.DB {
	var err error

	// open pool connection
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// config pool
	db.SetMaxOpenConns(25) // count open connections
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute) // security

	// security connections
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		err = db.PingContext(ctx)
		if err == nil {
			fmt.Println("Connected to database")
			return db
		}

		fmt.Printf("Failed to connect to database (attempt #%d): %v", i, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Failed to connect to database: %v", err)
	return nil

}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database")
		} else {
			fmt.Println("Pool Connection Closed")
		}
	}
}
