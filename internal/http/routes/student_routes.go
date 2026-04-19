package routes

import (
	"github.com/gorilla/mux"
	studentctrl "github.com/sahilchauhan0603/society/internal/service/student"
)

func registerStudentRoutes(api *mux.Router) {
	api.HandleFunc("/students", studentctrl.FetchAllStudents).Methods("GET")
	api.HandleFunc("/students/enroll/{enrollmentNo}", studentctrl.FetchStudent).Methods("GET")
	api.HandleFunc("/students/society/{societyID}", studentctrl.FetchStudentBySocietyID).Methods("GET")
	api.HandleFunc("/students/{enrollmentNo}/contributions", studentctrl.FetchContributions).Methods("GET")

	api.HandleFunc("/studentachievements", studentctrl.AddNewStudentAchievement).Methods("POST")
	api.HandleFunc("/studentachievements/{enrollmentNo}", studentctrl.UpdateStudentAchievement).Methods("PUT")
	api.HandleFunc("/studentachievements", studentctrl.FetchAllStudentAchievements).Methods("GET")
	api.HandleFunc("/studentachievements/{enrollmentNo}", studentctrl.RemoveStudentAchievement).Methods("DELETE")
	api.HandleFunc("/studentachievements/{enrollmentNo}", studentctrl.FetchStudentAchievements).Methods("GET")
	api.HandleFunc("/studentachievements/society/{societyID}", studentctrl.FetchStudentAchievementsSocietyID).Methods("GET")

	api.HandleFunc("/markings", studentctrl.AddNewMarking).Methods("POST")
	api.HandleFunc("/markings/{enrollmentNo}", studentctrl.UpdateMarking).Methods("PUT")
	api.HandleFunc("/markings", studentctrl.FetchAllMarkings).Methods("GET")
	api.HandleFunc("/markings/{societyID}", studentctrl.FetchMarkingSocietyID).Methods("GET")
	api.HandleFunc("/markings/{enrollmentNo}", studentctrl.RemoveMarking).Methods("DELETE")

	api.HandleFunc("/students", studentctrl.AddNewStudent).Methods("POST")
	api.HandleFunc("/students/{enrollmentNo}", studentctrl.UpdateStudent).Methods("PUT")
	api.HandleFunc("/admin/members", studentctrl.FetchAllStudentsAdmin).Methods("GET")
	api.HandleFunc("/admin/members/{societyID}", studentctrl.FetchStudentsSocietyAdmin).Methods("GET")
	api.HandleFunc("/students/{enrollmentNo}", studentctrl.RemoveStudent).Methods("DELETE")
}
