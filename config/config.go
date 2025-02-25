package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// initialize the database connection
func InitDB() {

	// Load environment variables from .env file
	if err:= godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}

	// Get database and connection details from env
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD") 
	dbHost := os.Getenv("DB_HOST")        
	dbPort := os.Getenv("DB_PORT")        
	dbName := os.Getenv("DB_NAME") 

	// Construct the database connection URL
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("mysql",url)
	if err != nil {
		log.Println("Error connecting DB", err)
		return
	}

	// Verify the connection to the database
	if err = DB.Ping(); err != nil {
		log.Println("Error pinging DB", err)
		return
	}
	
	log.Println("Database connected successfully..")
}