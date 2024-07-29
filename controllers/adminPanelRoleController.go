package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

func AddNewAdminRole(w http.ResponseWriter, r *http.Request) {
	var role models.AdminPanelRole
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AdminPanelRole{}) {
		if err := database.DB.AutoMigrate(&models.AdminPanelRole{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&role).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func FetchAllAdminRoles(w http.ResponseWriter, r *http.Request) {

	var roles []models.AdminPanelRole
	if err := database.DB.Find(&roles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

func FetchAdminRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	var role []models.AdminPanelRole
	if err := database.DB.First(&role, "username = ?", username).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func UpdateAdminRole(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]

    var role models.AdminPanelRole
    if err := database.DB.First(&role, "username = ?", username).Error; err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Save(&role).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(role)
}

func RemoveAdminRole(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    username := vars["username"]

    if err := database.DB.Delete(&models.AdminPanelRole{}, "username = ?", username).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Role successfully deleted"})
}
