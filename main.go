// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/handlers"
// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"

// 	database "github.com/sahilchauhan0603/society/config"
// 	routes "github.com/sahilchauhan0603/society/routes"
// )

// // // Define allowed origins
// // var allowedOrigins = []string{
// // 	"https://societymanagementfrontend-h3v3.onrender.com",
// // 	"https://societymanagementfrontend-h3v3.onrender.com/admin",
// // 	"http://localhost:8000",
// // }

// func main() {
// 	// Load environment variables
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Printf("Error loading .env file: %v", err)
// 	}

// 	// Initialize database connection
// 	database.DatabaseConnector()

// 	// Create a new router
// 	router := mux.NewRouter()
// 	routes.InitializeRoutes(router)

// 	// // CORS handler function
// 	// corsHandler := func(next http.Handler) http.Handler {
// 	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	// 		origin := r.Header.Get("Origin")
// 	// 		for _, allowedOrigin := range allowedOrigins {
// 	// 			if origin == allowedOrigin {
// 	// 				w.Header().Set("Access-Control-Allow-Origin", origin)
// 	// 				break
// 	// 			}
// 	// 		}
// 	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 	// 		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
// 	// 		w.Header().Set("Access-Control-Allow-Credentials", "true")

// 	// 		if r.Method == http.MethodOptions {
// 	// 			w.WriteHeader(http.StatusNoContent)
// 	// 			return
// 	// 		}

// 	// 		next.ServeHTTP(w, r)
// 	// 	})
// 	// }
// 	// Enable CORS
// 	cors := handlers.CORS(
// 		handlers.AllowedOrigins([]string{"https://societymanagementfrontend-h3v3.onrender.com"}), // Specific frontend origin
// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
// 		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
// 		handlers.AllowCredentials(), // Allow credentials
// 	)

// 	// // Apply CORS middleware
// 	// router.Use(corsHandler)

// 	// Set the port for the server
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080" // Default port if not specified
// 	}

// 	fmt.Printf("Server is running on port %s\n", port)
// 	log.Fatalf("Failed to start server: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), cors(router)))
// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	database "github.com/sahilchauhan0603/society/config"
	routes "github.com/sahilchauhan0603/society/routes"
)

// Define allowed origins
var allowedOrigins = []string{
	"https://societymanagementfrontend-h3v3.onrender.com",
	"http://localhost:8000",
	"http://localhost:5173",
}

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

	// CORS handler function
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowed := false

			// Check if the origin is in the allowed origins list
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					allowed = true
					break
				}
			}

			if allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Handle preflight OPTIONS request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// Proceed to the next handler for non-OPTIONS requests
			next.ServeHTTP(w, r)
		})
	}

	// Apply CORS middleware
	router.Use(corsHandler)

	// Set the port for the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatalf("Failed to start server: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
