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
        w.Header().Set("Access-Control-Allow-Origin", "https://societymanagementfrontend-h3v3.onrender.com") // Replace with your frontend's URL
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        w.Header().Set("Access-Control-Allow-Credentials", "true") // Allow credentials
        w.WriteHeader(http.StatusNoContent)
    })


	router.HandleFunc("/microsoftLogin", controllers.HandleMicrosoftLogin).Methods("GET")
	router.HandleFunc("/callback", controllers.HandleMicrosoftCallback).Methods("GET")

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/forgotPassword", controllers.SendEmail).Methods("POST")
	router.HandleFunc("/resetPassword", controllers.VerifyReset).Methods("POST")
	// registerHandler := http.HandlerFunc(controllers.Register)
	// router.Handle("/verifyOTP", middleware.OTPVerify(registerHandler)).Methods("POST")

	// router.HandleFunc("/delete-table/{table}/{column}", controllers.DeleteColumnHandler).Methods("DELETE")
	// router.HandleFunc("/delete-table/{table}", controllers.DeleteTableHandler).Methods("DELETE")


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
	r.HandleFunc("/roles", controllers.FetchAllRoles).Methods("GET")
	r.HandleFunc("/roles/name/{name}", controllers.FetchRole).Methods("GET")
	r.HandleFunc("/roles/society/{societyID}", controllers.FetchRoleSocietyID).Methods("GET")

	//MEMBERS
	r.HandleFunc("/students", controllers.FetchAllStudents).Methods("GET")
	r.HandleFunc("/students/enroll/{enrollmentNo}", controllers.FetchStudent).Methods("GET")
	r.HandleFunc("/students/society/{societyID}", controllers.FetchStudentBySocietyID).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}/contributions", controllers.FetchContributions).Methods("GET")

	//SOCIETY
	r.HandleFunc("/societies", controllers.FetchAllSocieties).Methods("GET")
	// r.HandleFunc("/societies/coordinator/{societyCoordinator}", controllers.FetchSocietyByCoordinator).Methods("GET")
	// r.HandleFunc("/societies/coordinator/{societyCoordinator}", controllers.RemoveSocietyByCoordinator).Methods("DELETE")
	//Fetch society by Society ID
	r.HandleFunc("/societies/{societyID}", controllers.FetchSocietyByID).Methods("GET")
	// get all members of a society
	r.HandleFunc("/societies/members/{societyID}", controllers.FetchStudentBySocietyID).Methods("GET")
	//query form
	r.HandleFunc("/societies/{societyID}/contact", controllers.SocietyQueryHandler).Methods("POST")
	r.HandleFunc("/createSociety", controllers.CreateSocietyHandler).Methods("POST")

	//society Achievements
	r.HandleFunc("/achievements", controllers.FetchAllAchievements).Methods("GET")
	r.HandleFunc("/achievements/{societyID}", controllers.FetchSocietyAchievementsSocietyID).Methods("GET")
	r.HandleFunc("/achievements/{societyID}", controllers.RemoveAchievement).Methods("DELETE")

	//SOCIETY EVENTS
	r.HandleFunc("/events", controllers.FetchAllEvents).Methods("GET")
	r.HandleFunc("/events/{eventID}", controllers.FetchEventByID).Methods("GET")
	r.HandleFunc("/events/society/{societyID}", controllers.FetchEventsBySocietyID).Methods("GET")
	r.HandleFunc("/events/society/{societyID}", controllers.RemoveEventsBySocietyID).Methods("DELETE")
	// r.HandleFunc("/events/{societyID}/{eventID}", controllers.RegisterForEvent).Methods("POST")
	r.HandleFunc("/registerForEvent", controllers.RegisterForEvent).Methods("POST")

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
	r.HandleFunc("/testimonials", controllers.FetchAllTestimonials).Methods("GET")
	r.HandleFunc("/testimonials/{societyID}", controllers.RemoveTestimonialSocietyID).Methods("DELETE")
	r.HandleFunc("/testimonials/{enrollmentNo}", controllers.FetchTestimonialByID).Methods("GET")
	r.HandleFunc("/testimonials/society/{societyID}", controllers.FetchTestimonialBySocietyID).Methods("GET")

	// Coordinator
	r.HandleFunc("/coordinator", controllers.FetchAllCoordinators).Methods("GET")
	r.HandleFunc("/coordinator/{societyID}", controllers.FetchCoordinatorByID).Methods("GET")
	r.HandleFunc("/coordinator/{coordinatorID}", controllers.FetchCoordinatorByCoordID).Methods("GET")

	//Gallery
	r.HandleFunc("/galleries", controllers.FetchAllGalleries).Methods("GET")
	r.HandleFunc("/galleries/{society_id}", controllers.FetchGallerySociety).Methods("GET")

	//News
	r.HandleFunc("/news", controllers.FetchAllNews).Methods("GET")
	r.HandleFunc("/news/{society_id}", controllers.FetchNews).Methods("GET")

	r.HandleFunc("/contact", controllers.ContactUSHandler).Methods("POST")
	r.HandleFunc("/feedback", controllers.FeedbackHandler).Methods("POST")
	r.HandleFunc("/becomeMember", controllers.BecomeMemberHandler).Methods("POST")



	// AdminRole routes
	router.HandleFunc("/adminlogin", controllers.AdminLogin).Methods("POST")
	adminRouter := router.PathPrefix("/api/v1/admin").Subrouter()
	adminRouter.Use(middleware.AdminMiddleware)
	adminRouter.HandleFunc("/adminrole", controllers.AddNewAdminRole).Methods("POST")
	r.HandleFunc("/adminroles", controllers.FetchAllAdminRoles).Methods("GET")
	r.HandleFunc("/adminrole/{username}", controllers.FetchAdminRole).Methods("GET")
	r.HandleFunc("/adminrole/{username}", controllers.UpdateAdminRole).Methods("PUT")
	adminRouter.HandleFunc("/adminrole/{username}", controllers.RemoveAdminRole).Methods("DELETE")


	//ADMIN PANEL ROUTES
	r.HandleFunc("/admin/home/news", controllers.FetchAllNewsAdminHome).Methods("GET")

    
	r.HandleFunc("/news", controllers.AddNewNews).Methods("POST")
	r.HandleFunc("/admin/news/{society_id}", controllers.FetchNewsAdminNews).Methods("GET")
	r.HandleFunc("/admin/news", controllers.FetchAllNewsAdminNews).Methods("GET")
	r.HandleFunc("/news/{newsID}", controllers.UpdateNews).Methods("PUT")
	r.HandleFunc("/news/{newsID}", controllers.RemoveNews).Methods("DELETE")


    r.HandleFunc("/students", controllers.AddNewStudent).Methods("POST")
	r.HandleFunc("/students/{enrollmentNo}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/admin/members", controllers.FetchAllStudentsAdmin).Methods("GET")
	r.HandleFunc("/admin/members/{societyID}", controllers.FetchStudentsSocietyAdmin).Methods("GET")
	r.HandleFunc("/students/{enrollmentNo}", controllers.RemoveStudent).Methods("DELETE")


	r.HandleFunc("/coordinator", controllers.AddNewCoordinator).Methods("POST")
	r.HandleFunc("/coordinator/{coordinatorID}", controllers.UpdateCoordinator).Methods("PUT")
	r.HandleFunc("/admin/coordinator", controllers.FetchAllCoordinatorsAdmin).Methods("GET")
	r.HandleFunc("/admin/coordinator/{societyID}", controllers.FetchCoordinatorAdminByID).Methods("GET")
	r.HandleFunc("/coordinator/{coordinatorID}", controllers.RemoveCoordinator).Methods("DELETE")


	r.HandleFunc("/events", controllers.AddNewEvent).Methods("POST")
	r.HandleFunc("/events/{eventID}", controllers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/admin/events", controllers.FetchAllAdminEvents).Methods("GET")
	r.HandleFunc("/admin/events/{societyID}", controllers.FetchAllAdminEventsSociety).Methods("GET")
	r.HandleFunc("/events/{eventID}", controllers.RemoveEvent).Methods("DELETE")


	r.HandleFunc("/societies", controllers.AddNewSociety).Methods("POST")
	r.HandleFunc("/societies/{societyID}", controllers.UpdateSociety).Methods("PUT")
	r.HandleFunc("/admin/societies", controllers.FetchAllSocietiesAdmin).Methods("GET")
	r.HandleFunc("/admin/societies/{societyID}", controllers.FetchSocietyAdmin).Methods("GET")
	r.HandleFunc("/societies/{societyID}", controllers.RemoveSocietyByID).Methods("DELETE")


	r.HandleFunc("/testimonials", controllers.AddNewTestimonial).Methods("POST")
	r.HandleFunc("/testimonials/{testimonialID}", controllers.UpdateTestimonial).Methods("PUT")
	r.HandleFunc("/admin/testimonials", controllers.FetchAllTestimonialsAdmin).Methods("GET")
	r.HandleFunc("/admin/testimonials/{societyID}", controllers.FetchAllTestimonialsSocietyAdmin).Methods("GET")
	r.HandleFunc("/testimonials/{testimonialID}", controllers.RemoveTestimonial).Methods("DELETE")


	r.HandleFunc("/achievements", controllers.AddNewAchievement).Methods("POST")
	r.HandleFunc("/achievements/{achievementID}", controllers.UpdateAchievement).Methods("PUT")
	r.HandleFunc("/admin/achievements",controllers.FetchAllAchievements).Methods("GET")
	r.HandleFunc("/admin/achievements/{societyID}",controllers.FetchSocietyAchievementsSocietyID).Methods("GET")
	r.HandleFunc("/achievements/{achievementID}", controllers.RemoveAchievementAchievementID).Methods("DELETE")


    r.HandleFunc("/galleries", controllers.AddNewGallery).Methods("POST")
	r.HandleFunc("/admin/gallery", controllers.FetchAllGalleries).Methods("GET")
	r.HandleFunc("/admin/gallery/{society_id}", controllers.FetchGallery).Methods("GET")
	r.HandleFunc("/galleries/{societyID}", controllers.UpdateGallery).Methods("PUT")
	r.HandleFunc("/galleries/{societyID}", controllers.RemoveGallery).Methods("DELETE")
    

	r.HandleFunc("/roles", controllers.AddNewRole).Methods("POST")
	r.HandleFunc("/roles/{roleID}", controllers.UpdateRole).Methods("PUT")
	r.HandleFunc("/admin/roles",controllers.FetchAllRolesAdmin).Methods("GET")
	r.HandleFunc("/admin/roles/{societyID}",controllers.FetchAllRolesSocietyAdmin).Methods("GET")
	r.HandleFunc("/roles/{roleID}", controllers.RemoveRole).Methods("DELETE")
}
