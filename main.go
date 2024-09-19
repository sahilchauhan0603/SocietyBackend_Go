package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	database "github.com/sahilchauhan0603/society/config"
	routes "github.com/sahilchauhan0603/society/routes"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	database.DatabaseConnector()

	// Create a new router
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	// Enable CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"https://societymanagementfrontend-h3v3.onrender.com"}), // Specific frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowCredentials(), // Allow credentials
	)

	// Set the port for the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatalf("Failed to start server: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), cors(router)))
}
