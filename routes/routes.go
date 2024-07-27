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

	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	registerHandler := http.HandlerFunc(controllers.Register)
	router.Handle("/verifyOTP", middleware.OTPVerify(registerHandler)).Methods("POST")

	router.HandleFunc("/delete-table/{table}", controllers.DeleteTableHandler).Methods("DELETE")

	// User routes
	r := router.PathPrefix("/api/v1").Subrouter()
	// r.Use(middleware.JWTVerify)
	
    //USER
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/userID/{id}", controllers.GetUserID).Methods("GET")
	r.HandleFunc("/user/society/{societyID}", controllers.FetchUsersSocietyID).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

    
	//SOCIETY ROLES
	r.HandleFunc("/roles", controllers.AddNewRole).Methods("POST")
	r.HandleFunc("/roles/{id}", controllers.UpdateRole).Methods("PUT")
	r.HandleFunc("/roles", controllers.FetchAllRoles).Methods("GET")
	r.HandleFunc("/roles/name/{name}", controllers.FetchRole).Methods("GET")
	r.HandleFunc("/roles/society/{societyID}", controllers.FetchRoleSocietyID).Methods("GET")
	r.HandleFunc("/roles/{roleID}", controllers.RemoveRole).Methods("DELETE")

    
	//MEMBERS
	r.HandleFunc("/students", controllers.AddNewStudent).Methods("POST")
	r.HandleFunc("/students/{enrollmentNo}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students", controllers.FetchAllStudents).Methods("GET")
	r.HandleFunc("/students/enroll/{enrollmentNo}", controllers.FetchStudent).Methods("GET")
	r.HandleFunc("/students/society/{societyID}", controllers.FetchStudentBySocietyID).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}/contributions", controllers.FetchContributions).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}", controllers.RemoveStudent).Methods("DELETE")

    
	//SOCIETY
	r.HandleFunc("/societies", controllers.AddNewSociety).Methods("POST")
	r.HandleFunc("/societies/{societyID}", controllers.UpdateSociety).Methods("PUT")
	r.HandleFunc("/societies", controllers.FetchAllSocieties).Methods("GET")
	r.HandleFunc("/societies/coordinator/{societyCoordinator}", controllers.FetchSocietyByCoordinator).Methods("GET")
	// r.HandleFunc("/societies/{societyID}", controllers.FetchSocietyByID).Methods("GET")
	r.HandleFunc("/societies/{societyID}", controllers.RemoveSocietyByID).Methods("DELETE")
	r.HandleFunc("/societies/coordinator/{societyCoordinator}", controllers.RemoveSocietyByCoordinator).Methods("DELETE")
	//Fetch society by Society ID
	r.HandleFunc("/societies/{societyID}", controllers.FetchSocietyByID).Methods("GET")
	// get all members of a society
	r.HandleFunc("/societies/members/{societyID}",controllers.FetchStudentBySocietyID).Methods("GET")


    //society Achievements
	r.HandleFunc("/achievements", controllers.AddNewAchievement).Methods("POST")
	r.HandleFunc("/achievements/{societyID}", controllers.UpdateAchievement).Methods("PUT")
	r.HandleFunc("/achievements", controllers.FetchAllAchievements).Methods("GET")
	r.HandleFunc("/achievements/{societyID}", controllers.FetchSocietyAchievementsSocietyID).Methods("GET")
	r.HandleFunc("/achievements/{societyID}", controllers.RemoveAchievement).Methods("DELETE")


	//SOCIETY EVENTS
	r.HandleFunc("/events", controllers.AddNewEvent).Methods("POST")
	r.HandleFunc("/events/{societyID}", controllers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events", controllers.FetchAllEvents).Methods("GET")
	r.HandleFunc("/events/{eventID}", controllers.FetchEventByID).Methods("GET")
	r.HandleFunc("/events/society/{societyID}", controllers.FetchEventsBySocietyID).Methods("GET")
	r.HandleFunc("/events/{eventID}", controllers.RemoveEvent).Methods("DELETE")
	r.HandleFunc("/events/society/{societyID}", controllers.RemoveEventsBySocietyID).Methods("DELETE")
	r.HandleFunc("/events/{societyID}/{eventID}", controllers.RegisterEventHandler).Methods("POST")

	// Student Achievements
	r.HandleFunc("/studentachievements", controllers.AddNewStudentAchievement).Methods("POST")
	r.HandleFunc("/studentachievements/{enrollmentNo}", controllers.UpdateStudentAchievement).Methods("PUT")
	r.HandleFunc("/studentachievements", controllers.FetchAllStudentAchievements).Methods("GET")
	r.HandleFunc("/studentachievements/{enrollmentNo}", controllers.RemoveStudentAchievement).Methods("DELETE")
	r.HandleFunc("/studentachievements/{enrollmentNo}", controllers.FetchStudentAchievements).Methods("GET")
	r.HandleFunc("/studentachievements/society/{societyID}", controllers.FetchStudentAchievementsSocietyID).Methods("GET")

	// Student Markings
	r.HandleFunc("/markings", controllers.AddNewMarking).Methods("POST")
	r.HandleFunc("/markings/{enrollmentNo}", controllers.UpdateMarking).Methods("PUT")
	r.HandleFunc("/markings", controllers.FetchAllMarkings).Methods("GET")
	r.HandleFunc("/markings/{societyID}", controllers.FetchMarkingSocietyID).Methods("GET")
	r.HandleFunc("/markings/{enrollmentNo}", controllers.RemoveMarking).Methods("DELETE")


	// Testimonials
	r.HandleFunc("/testimonials", controllers.AddNewTestimonial).Methods("POST")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.UpdateTestimonial).Methods("PUT")
	r.HandleFunc("/testimonials", controllers.FetchAllTestimonials).Methods("GET")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.RemoveTestimonial).Methods("DELETE")
	r.HandleFunc("/testimonials/{societyID}", controllers.RemoveTestimonialSocietyID).Methods("DELETE")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.FetchTestimonialByID).Methods("GET")
	r.HandleFunc("/testimonials/society/{societyID}", controllers.FetchTestimonialBySocietyID).Methods("GET")


	// Coordinator
	r.HandleFunc("/coordinator", controllers.AddNewCoordinator).Methods("POST")
	r.HandleFunc("/coordinator/{societyID}", controllers.UpdateCoordinator).Methods("PUT")
	r.HandleFunc("/coordinator", controllers.FetchAllCoordinators).Methods("GET")
	r.HandleFunc("/coordinator/{societyID}", controllers.RemoveCoordinator).Methods("DELETE")
	r.HandleFunc("/coordinator/{societyID}", controllers.FetchCoordinatorByID).Methods("GET")


	//Gallery
	r.HandleFunc("/galleries", controllers.AddNewGallery).Methods("POST")
	r.HandleFunc("/galleries", controllers.FetchAllGalleries).Methods("GET")
	r.HandleFunc("/galleries/{society_id}", controllers.FetchGallery).Methods("GET")
	r.HandleFunc("/galleries/{societyID}", controllers.UpdateGallery).Methods("PUT")
	r.HandleFunc("/galleries/{societyID}", controllers.RemoveGallery).Methods("DELETE")


	//News
	r.HandleFunc("/news", controllers.AddNewNews).Methods("POST")
	r.HandleFunc("/news", controllers.FetchAllNews).Methods("GET")
	r.HandleFunc("/news/{society_id}", controllers.FetchNews).Methods("GET")
	r.HandleFunc("/news/{societyID}", controllers.UpdateNews).Methods("PUT")
	r.HandleFunc("/news/{societyID}", controllers.RemoveNews).Methods("DELETE")


	r.HandleFunc("/contact",controllers.ContactUSHandler).Methods("POST")

}
