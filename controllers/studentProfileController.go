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
	id, err := strconv.Atoi(params["id"])
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stud)
}
func FetchAllStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.StudentProfile
	if err := database.DB.Preload("StudentAchievements").Order("user_id ASC").Find(&students).Error; err != nil {
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

	var student models.StudentProfile
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

	var student models.StudentProfile
	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).Select("student_contributions").First(&student).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student.StudentContributions)
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