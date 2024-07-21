package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/society/controllers"
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
	userRouter := router.PathPrefix("/api/v1").Subrouter()
	// uploaderRouter.Use(middleware.JWTVerify)
	userRouter.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/user", controllers.GetUser).Methods("GET")
	userRouter.HandleFunc("/user/{id}", controllers.GetUserID).Methods("GET")
	userRouter.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	roleRouter := router.PathPrefix("/api/v1").Subrouter()

	roleRouter.HandleFunc("/roles", controllers.AddNewRole).Methods("POST")
	roleRouter.HandleFunc("/roles/{id}", controllers.UpdateRole).Methods("PUT")
	roleRouter.HandleFunc("/roles", controllers.FetchAllRoles).Methods("GET")
	roleRouter.HandleFunc("/roles/{name}", controllers.FetchRole).Methods("GET")
	roleRouter.HandleFunc("/roles/{name}", controllers.RemoveRole).Methods("DELETE")
}
