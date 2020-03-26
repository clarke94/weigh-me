package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
)

// Connect will connect to the mongo database
func Connect() (*sql.DB, error) {
	// Get Environment variables
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// open the connection
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Print("Successfully connected")
	// return the connection
	return db, nil
}
