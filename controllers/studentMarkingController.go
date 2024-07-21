package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	"github.com/sahilchauhan0603/society/models"
)

func AddNewMarking(w http.ResponseWriter, r *http.Request) {
	var marking models.StudentMarking
	if err := json.NewDecoder(r.Body).Decode(&marking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.StudentMarking{}) {
		if err := database.DB.AutoMigrate(&models.StudentMarking{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&marking).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marking)
}
func UpdateMarking(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var marking models.StudentMarking
	if result := database.DB.First(&marking, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&marking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&marking)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(marking)
}
func FetchAllMarkings(w http.ResponseWriter, r *http.Request) {
	var markings []models.StudentMarking
	if err := database.DB.Order("enrollment_no ASC").Find(&markings).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(markings)
}
func RemoveMarking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo := vars["enrollmentNo"]

	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).Delete(&models.StudentMarking{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Marking successfully deleted"})
}
