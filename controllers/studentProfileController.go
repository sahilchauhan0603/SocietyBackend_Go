package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

func AddNewStudent(w http.ResponseWriter, r *http.Request) {
	var stud models.StudentProfile
	if err := json.NewDecoder(r.Body).Decode(&stud); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.StudentProfile{}) {
		if err := database.DB.AutoMigrate(&models.StudentProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&stud).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stud)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var stud models.StudentProfile
	if result := database.DB.First(&stud, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&stud); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&stud)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stud)
}

func FetchAllStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.StudentProfile
	if err := database.DB.Preload("StudentAchievements").Order("society_id ASC").Find(&students).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
func FetchStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo, err := strconv.ParseUint(vars["enrollmentNo"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var student []models.StudentProfile
	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).First(&student).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func FetchContributions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo, err := strconv.ParseUint(vars["enrollmentNo"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var student []models.StudentProfile
	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).Select("student_contributions").First(&student).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func RemoveStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo, err := strconv.ParseUint(vars["enrollmentNo"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).Delete(&models.StudentProfile{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Student successfully deleted"})
}

func FetchStudentBySocietyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID, err := strconv.ParseUint(vars["societyID"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var student []models.StudentProfile
	if err := database.DB.Where("society_id = ?", societyID).Find(&student).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// ADMIN PANEL
func FetchAllStudentsAdmin(w http.ResponseWriter, r *http.Request) {

	var tempMembers []struct {
		ProfilePicture       string
		FirstName            string
		LastName             string
		Branch               string
		BatchYear            int
		EnrollmentNo         uint
		Email                string
		// StudentContributions string
	}
	if err := database.DB.Model(&models.StudentProfile{}).
		Select("profile_picture, first_name, last_name, branch, batch_year, enrollment_no, email").
		Find(&tempMembers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempMembers)
}

func FetchStudentsSocietyAdmin(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID, err := strconv.ParseUint(vars["societyID"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tempMembers []struct {
		EnrollmentNo         uint
		ProfilePicture       string
		FirstName            string
		LastName             string
		Branch               string
		BatchYear            int
		MobileNo             string
		Email                string
		StudentContributions string
	}
	if err := database.DB.Model(&models.StudentProfile{}).
		Select("profile_picture, first_name, last_name, branch, batch_year, mobile_no, email, enrollment_no, student_contributions").
		Where("society_id = ?", societyID).
		Find(&tempMembers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempMembers)
}
