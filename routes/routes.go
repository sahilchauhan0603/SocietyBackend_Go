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
	r := router.PathPrefix("/api/v1").Subrouter()
	// r.Use(middleware.JWTVerify)
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserID).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/roles", controllers.AddNewRole).Methods("POST")
	r.HandleFunc("/roles/{id}", controllers.UpdateRole).Methods("PUT")
	r.HandleFunc("/roles", controllers.FetchAllRoles).Methods("GET")
	r.HandleFunc("/roles/{name}", controllers.FetchRole).Methods("GET")
	r.HandleFunc("/roles/{name}", controllers.RemoveRole).Methods("DELETE")


	r.HandleFunc("/students", controllers.AddNewStudent).Methods("POST")
	r.HandleFunc("/students/{enrollmentNo}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students", controllers.FetchAllStudents).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}", controllers.FetchStudent).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}/contributions", controllers.FetchContributions).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}", controllers.RemoveStudent).Methods("DELETE")

	r.HandleFunc("/societies", controllers.AddNewSociety).Methods("POST")
	r.HandleFunc("/societies/{societyID}", controllers.UpdateSociety).Methods("PUT")
	r.HandleFunc("/societies", controllers.FetchAllSocieties).Methods("GET")
	r.HandleFunc("/societies/coordinator/{societyCoordinator}", controllers.FetchSocietyByCoordinator).Methods("GET")
	r.HandleFunc("/societies/{societyID}", controllers.FetchSocietyByID).Methods("GET")
	r.HandleFunc("/societies/{societyID}", controllers.RemoveSocietyByID).Methods("DELETE")
	r.HandleFunc("/societies/coordinator/{societyCoordinator}", controllers.RemoveSocietyByCoordinator).Methods("DELETE")


	r.HandleFunc("/achievements", controllers.AddNewAchievement).Methods("POST")
	r.HandleFunc("/achievements/{societyID}", controllers.UpdateAchievement).Methods("PUT")
	r.HandleFunc("/achievements", controllers.FetchAllAchievements).Methods("GET")
	r.HandleFunc("/achievements/{societyID}", controllers.RemoveAchievement).Methods("DELETE")

	r.HandleFunc("/events", controllers.AddNewEvent).Methods("POST")
	r.HandleFunc("/events/{eventID}", controllers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events", controllers.FetchAllEvents).Methods("GET")
	r.HandleFunc("/events/{eventID}", controllers.FetchEventByID).Methods("GET")
	r.HandleFunc("/events/society/{societyID}", controllers.FetchEventsBySocietyID).Methods("GET")
	r.HandleFunc("/events/{eventID}", controllers.RemoveEvent).Methods("DELETE")
	r.HandleFunc("/events/society/{societyID}", controllers.RemoveEventsBySocietyID).Methods("DELETE")

	// Student Achievements
	r.HandleFunc("/achievements", controllers.AddNewStudentAchievement).Methods("POST")
	r.HandleFunc("/achievements/{enrollmentNo}", controllers.UpdateStudentAchievement).Methods("PUT")
	r.HandleFunc("/achievements", controllers.FetchAllStudentAchievements).Methods("GET")
	r.HandleFunc("/achievements/{enrollmentNo}", controllers.RemoveStudentAchievement).Methods("DELETE")
	r.HandleFunc("/achievements/{enrollmentNo}", controllers.FetchStudentAchievements).Methods("GET")

	// Student Markings
	r.HandleFunc("/markings", controllers.AddNewMarking).Methods("POST")
	r.HandleFunc("/markings/{enrollmentNo}", controllers.UpdateMarking).Methods("PUT")
	r.HandleFunc("/markings", controllers.FetchAllMarkings).Methods("GET")
	r.HandleFunc("/markings/{enrollmentNo}", controllers.RemoveMarking).Methods("DELETE")

	// Testimonials
	r.HandleFunc("/testimonials", controllers.AddNewTestimonial).Methods("POST")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.UpdateTestimonial).Methods("PUT")
	r.HandleFunc("/testimonials", controllers.FetchAllTestimonials).Methods("GET")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.RemoveTestimonial).Methods("DELETE")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.FetchTestimonialByID).Methods("GET")
}
