package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/society/controllers"
	"github.com/sahilchauhan0603/society/middleware"
)

func InitializeRoutes(router *mux.Router) {

	// Handle preflight requests for the /api/v1 endpoints
	router.PathPrefix("/api/v1").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.WriteHeader(http.StatusNoContent)
	})

	router.HandleFunc("/login", controllers.HandleMicrosoftLogin).Methods("GET")
	router.HandleFunc("/callback", controllers.HandleMicrosoftCallback).Methods("GET")

	// User routes
	uploaderRouter := router.PathPrefix("/api/v1").Subrouter()
	uploaderRouter.Use(middleware.JWTVerify)
	uploaderRouter.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	uploaderRouter.HandleFunc("/user", controllers.GetUser).Methods("GET")
	uploaderRouter.HandleFunc("/user/{id}", controllers.GetUserID).Methods("GET")
	uploaderRouter.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	uploaderRouter.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
