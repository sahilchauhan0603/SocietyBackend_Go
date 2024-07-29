package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

type tempo struct {
	SocietyID              uint
	CoordinatorName        string
	CoordinatorDesignation string
	CoordinatorEmail       string
	CoordinatorDetails     string
	SocietyName            string
	Image                  string
}

func AddNewCoordinator(w http.ResponseWriter, r *http.Request) {

	var coordinator models.SocietyCoordinator
	if err := json.NewDecoder(r.Body).Decode(&coordinator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyCoordinator{}) {
		if err := database.DB.AutoMigrate(&models.SocietyCoordinator{}); err != nil {
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

	var coordinator models.SocietyCoordinator
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
	if err := database.DB.Table("society_coordinators").
		Select("society_coordinators.society_id, society_coordinators.coordinator_details, society_coordinators.coordinator_name, society_coordinators.coordinator_email, society_coordinators.coordinator_designation, society_coordinators.image, society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_coordinators.society_id").
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

	var info []tempo
	if err := database.DB.Table("society_coordinators").
	    Select("society_coordinators.society_id, society_coordinators.coordinator_details, society_coordinators.coordinator_name, society_coordinators.coordinator_email, society_coordinators.coordinator_designation, society_coordinators.image, society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_coordinators.society_id").Where("society_coordinators.society_id = ?", societyID).
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

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.SocietyCoordinator{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Coordinator successfully deleted"})
}
