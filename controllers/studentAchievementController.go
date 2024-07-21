package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	"github.com/sahilchauhan0603/society/models"
)

func AddNewStudentAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.StudentAchievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.StudentAchievement{}) {
		if err := database.DB.AutoMigrate(&models.StudentAchievement{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&achievement).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievement)
}
func UpdateStudentAchievement(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var achievement models.StudentAchievement
	if result := database.DB.First(&achievement, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&achievement)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(achievement)
}
func FetchAllStudentAchievements(w http.ResponseWriter, r *http.Request) {
	var achievements []models.StudentAchievement
	if err := database.DB.Order("achievement_id ASC").Find(&achievements).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievements)
}
func FetchStudentAchievements(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo := vars["enrollmentNo"]

	var achievement models.StudentAchievement
	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).First(&achievement).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievement)
}
func RemoveStudentAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo := vars["enrollmentNo"]

	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).Delete(&models.StudentAchievement{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Achievement successfully deleted"})
}
