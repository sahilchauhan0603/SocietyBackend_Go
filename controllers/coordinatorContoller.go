package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

type tempo struct {
	SocietyID          uint
	CoordinatorDetails string
	SocietyType        string
	SocietyName        string
	SocietyHead        string
	DateOfRegistration time.Time
	SocietyDescription string
}

func AddNewCoordinator(w http.ResponseWriter, r *http.Request) {

	var coordinator models.Coordinator
	if err := json.NewDecoder(r.Body).Decode(&coordinator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Coordinator{}) {
		if err := database.DB.AutoMigrate(&models.Coordinator{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&coordinator).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coordinator)
}

func UpdateCoordinator(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["societyID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var coordinator models.Coordinator
	if result := database.DB.First(&coordinator, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&coordinator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&coordinator)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coordinator)
}

func FetchAllCoordinators(w http.ResponseWriter, r *http.Request) {

	var data []tempo
	if err := database.DB.Table("coordinators").
		Select("coordinators.society_id, coordinators.coordinator_details, society_profiles.society_type, society_profiles.society_name, society_profiles.society_head, society_profiles.date_of_registration, society_profiles.society_description").
		Joins("JOIN society_profiles ON society_profiles.society_id = coordinators.society_id").
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func FetchCoordinatorByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	var info tempo
	if err := database.DB.Table("coordinators").
	    Select("coordinators.society_id, coordinators.coordinator_details, society_profiles.society_type, society_profiles.society_name, society_profiles.society_head, society_profiles.date_of_registration, society_profiles.society_description").
		Joins("JOIN society_profiles ON society_profiles.society_id = coordinators.society_id").Where("coordinators.society_id = ?", societyID).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func RemoveCoordinator(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.Coordinator{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Coordinator successfully deleted"})
}