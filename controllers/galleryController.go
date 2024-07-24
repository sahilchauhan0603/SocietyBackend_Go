package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

func AddNewGallery(w http.ResponseWriter, r *http.Request) {

	var gallery models.SocietyGallery
	if err := json.NewDecoder(r.Body).Decode(&gallery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyGallery{}) {
		if err := database.DB.AutoMigrate(&models.SocietyGallery{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := database.DB.Create(&gallery).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gallery)
}


func FetchAllGalleries(w http.ResponseWriter, r *http.Request) {

	var galleries []models.SocietyGallery
	if result := database.DB.Find(&galleries); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(galleries)
}

func FetchGallery(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["society_id"])
	if err != nil {
		http.Error(w, "Invalid Society ID", http.StatusBadRequest)
		return
	}

	var gallery models.SocietyGallery
	if err := database.DB.Where("society_id = ?", societyID).First(&gallery).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gallery)
}

func UpdateGallery(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["societyID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var gallery models.SocietyGallery
	if result := database.DB.First(&gallery, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&gallery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&gallery)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gallery)
}

func RemoveGallery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["societyID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.SocietyGallery{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Society Gallery successfully deleted"})
}
