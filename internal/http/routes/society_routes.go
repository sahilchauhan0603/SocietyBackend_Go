package routes

import (
	"github.com/gorilla/mux"
	societyctrl "github.com/sahilchauhan0603/society/internal/http/controllers/society"
	studentctrl "github.com/sahilchauhan0603/society/internal/http/controllers/student"
)

func registerSocietyRoutes(api *mux.Router) {
	api.HandleFunc("/user", societyctrl.CreateUser).Methods("POST")
	api.HandleFunc("/user", societyctrl.GetUser).Methods("GET")
	api.HandleFunc("/user/userID/{id}", societyctrl.GetUserID).Methods("GET")
	api.HandleFunc("/user/society/{societyID}", societyctrl.FetchUsersSocietyID).Methods("GET")
	api.HandleFunc("/user/{id}", societyctrl.UpdateUser).Methods("PUT")
	api.HandleFunc("/user/{id}", societyctrl.DeleteUser).Methods("DELETE")

	api.HandleFunc("/roles", societyctrl.FetchAllRoles).Methods("GET")
	api.HandleFunc("/roles/name/{name}", societyctrl.FetchRole).Methods("GET")
	api.HandleFunc("/roles/society/{societyID}", societyctrl.FetchRoleSocietyID).Methods("GET")

	api.HandleFunc("/societies", societyctrl.FetchAllSocieties).Methods("GET")
	api.HandleFunc("/societies/{societyID}", societyctrl.FetchSocietyByID).Methods("GET")
	api.HandleFunc("/societies/members/{societyID}", studentctrl.FetchStudentBySocietyID).Methods("GET")
	api.HandleFunc("/societies/{societyID}/contact", societyctrl.SocietyQueryHandler).Methods("POST")
	api.HandleFunc("/createSociety", societyctrl.CreateSocietyHandler).Methods("POST")

	api.HandleFunc("/achievements", societyctrl.FetchAllAchievements).Methods("GET")
	api.HandleFunc("/achievements/{societyID}", societyctrl.FetchSocietyAchievementsSocietyID).Methods("GET")
	api.HandleFunc("/achievements/{societyID}", societyctrl.RemoveAchievement).Methods("DELETE")

	api.HandleFunc("/events", societyctrl.FetchAllEvents).Methods("GET")
	api.HandleFunc("/events/{eventID}", societyctrl.FetchEventByID).Methods("GET")
	api.HandleFunc("/events/society/{societyID}", societyctrl.FetchEventsBySocietyID).Methods("GET")
	api.HandleFunc("/events/society/{societyID}", societyctrl.RemoveEventsBySocietyID).Methods("DELETE")
	api.HandleFunc("/registerForEvent", societyctrl.RegisterForEvent).Methods("POST")

	api.HandleFunc("/coordinator", societyctrl.FetchAllCoordinators).Methods("GET")
	api.HandleFunc("/coordinator/{societyID}", societyctrl.FetchCoordinatorByID).Methods("GET")
	api.HandleFunc("/coordinator/{coordinatorID}", societyctrl.FetchCoordinatorByCoordID).Methods("GET")

	api.HandleFunc("/coordinator", societyctrl.AddNewCoordinator).Methods("POST")
	api.HandleFunc("/coordinator/{coordinatorID}", societyctrl.UpdateCoordinator).Methods("PUT")
	api.HandleFunc("/admin/coordinator", societyctrl.FetchAllCoordinatorsAdmin).Methods("GET")
	api.HandleFunc("/admin/coordinator/{societyID}", societyctrl.FetchCoordinatorAdminByID).Methods("GET")
	api.HandleFunc("/coordinator/{coordinatorID}", societyctrl.RemoveCoordinator).Methods("DELETE")

	api.HandleFunc("/events", societyctrl.AddNewEvent).Methods("POST")
	api.HandleFunc("/events/{eventID}", societyctrl.UpdateEvent).Methods("PUT")
	api.HandleFunc("/admin/events", societyctrl.FetchAllAdminEvents).Methods("GET")
	api.HandleFunc("/admin/events/{societyID}", societyctrl.FetchAllAdminEventsSociety).Methods("GET")
	api.HandleFunc("/events/{eventID}", societyctrl.RemoveEvent).Methods("DELETE")

	api.HandleFunc("/societies", societyctrl.AddNewSociety).Methods("POST")
	api.HandleFunc("/societies/{societyID}", societyctrl.UpdateSociety).Methods("PUT")
	api.HandleFunc("/admin/societies", societyctrl.FetchAllSocietiesAdmin).Methods("GET")
	api.HandleFunc("/admin/societies/{societyID}", societyctrl.FetchSocietyAdmin).Methods("GET")
	api.HandleFunc("/societies/{societyID}", societyctrl.RemoveSocietyByID).Methods("DELETE")

	api.HandleFunc("/achievements", societyctrl.AddNewAchievement).Methods("POST")
	api.HandleFunc("/achievements/{achievementID}", societyctrl.UpdateAchievement).Methods("PUT")
	api.HandleFunc("/admin/achievements", societyctrl.FetchAllAchievements).Methods("GET")
	api.HandleFunc("/admin/achievements/{societyID}", societyctrl.FetchSocietyAchievementsSocietyID).Methods("GET")
	api.HandleFunc("/achievements/{achievementID}", societyctrl.RemoveAchievementAchievementID).Methods("DELETE")

	api.HandleFunc("/roles", societyctrl.AddNewRole).Methods("POST")
	api.HandleFunc("/roles/{roleID}", societyctrl.UpdateRole).Methods("PUT")
	api.HandleFunc("/admin/roles", societyctrl.FetchAllRolesAdmin).Methods("GET")
	api.HandleFunc("/admin/roles/{societyID}", societyctrl.FetchAllRolesSocietyAdmin).Methods("GET")
	api.HandleFunc("/roles/{roleID}", societyctrl.RemoveRole).Methods("DELETE")
}
