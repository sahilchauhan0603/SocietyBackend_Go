package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	"github.com/sahilchauhan0603/society/models"
)

func AddNewAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.SocietyAchievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyAchievement{}) {
		if err := database.DB.AutoMigrate(&models.SocietyAchievement{}); err != nil {
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
func UpdateAchievement(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["achievementID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var achievement models.SocietyAchievement
	if result := database.DB.First(&achievement, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&achievement)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(achievement)
}

func FetchAllAchievements(w http.ResponseWriter, r *http.Request) {
	
	var achievements []models.SocietyAchievement
	if err := database.DB.Order("society_achievement_id ASC").Find(&achievements).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievements)
}

func FetchSocietyAchievementsSocietyID(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	societyID := vars["societyID"]

	var achievement []models.SocietyAchievement
	if err := database.DB.Where("society_id = ?", societyID).First(&achievement).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievement)
}

func RemoveAchievement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["societyID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.SocietyAchievement{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Achievement successfully deleted"})
}
func RemoveAchievementAchievementID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	achievementID, err := strconv.Atoi(vars["achievementID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("society_achievement_id = ?", achievementID).Delete(&models.SocietyAchievement{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Achievement successfully deleted"})
}

