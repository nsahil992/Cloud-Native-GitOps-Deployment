package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"calico-go-project/database"
	"calico-go-project/handlers"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Initialize database connection
	err = database.InitDB()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	defer database.CloseDB()

	// Set up routes
	router := handlers.SetupRoutes()

	// Start server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
