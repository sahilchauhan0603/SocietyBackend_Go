package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

func AddNewSociety(w http.ResponseWriter, r *http.Request) {
	var society models.SocietyProfile
	if err := json.NewDecoder(r.Body).Decode(&society); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyProfile{}) {
		if err := database.DB.AutoMigrate(&models.SocietyProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&society).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(society)
}
func UpdateSociety(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var society models.SocietyProfile
	if result := database.DB.First(&society, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&society); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&society)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(society)
}
func FetchAllSocieties(w http.ResponseWriter, r *http.Request) {
	var societies []models.SocietyProfile
	if err := database.DB.Order("society_id ASC").Find(&societies).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(societies)
}
func FetchSocietyByCoordinator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyCoordinator := vars["societyCoordinator"]

	var society models.SocietyProfile
	if err := database.DB.Where("society_coordinator = ?", societyCoordinator).First(&society).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(society)
}
func FetchSocietyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["societyID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var society models.SocietyProfile
	if err := database.DB.Where("society_id = ?", societyID).First(&society).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(society)
}
func RemoveSocietyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["societyID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.SocietyProfile{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Society successfully deleted"})
}
func RemoveSocietyByCoordinator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyCoordinator := vars["societyCoordinator"]

	if err := database.DB.Where("society_coordinator = ?", societyCoordinator).Delete(&models.SocietyProfile{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Society successfully deleted"})
}
