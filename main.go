package main

import (
	"Codimite_Assignment/config"
	"Codimite_Assignment/routers"
	"log"
	"net/http"

	// "net/http"
	"os"
)

func main() {
	
	// Set up the routes
	r := routers.SetRoutes()

	// Initialize the database connection
	config.InitDB()

	// Define the port and the server
	port := ":8080"
	log.Println("Server is starting on port 8080....")

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Println("Server is not running",err)
		os.Exit(1)
	}
	
}