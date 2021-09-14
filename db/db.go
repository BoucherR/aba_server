package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//DB instance
var DB *sql.DB

//Connect to db
func Connect() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, _ := sql.Open("postgres", dbinfo)
	err = db.Ping()
	if err != nil {
		// log.Printf("Attempted to connect with: user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
		log.Fatalf("ERROR: Attempted to connect to DB with: user=%s password=%s dbname=%s sslmode=disable. ERR=%v", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), err)
	}
	DB = db
	// Create "users" table if it doesnt exist
	CreateUsersTable()
}
