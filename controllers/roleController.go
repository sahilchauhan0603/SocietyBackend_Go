package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
)

func AddNewRole(w http.ResponseWriter, r *http.Request) {

	var role models.SocietyRole
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyRole{}) {
		if err := database.DB.AutoMigrate(&models.SocietyRole{}); err != nil {
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

func FetchAllRoles(w http.ResponseWriter, r *http.Request) {

	var roles []struct {
		RoleID          uint
		RoleType        string
		Rolename        string
		RoleDescription string
	}
	if err := database.DB.Model(&models.SocietyRole{}).
		Select("role_id, role_type, rolename, role_description").
		Order("role_id ASC").
		Find(&roles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

func FetchRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rolename := vars["name"]

	var role []models.SocietyRole
	if err := database.DB.Where("rolename = ?", rolename).First(&role).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func FetchRoleSocietyID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	type roles struct {
		RoleID           int64
		RoleType         string
		Rolename         string
		RoleDescription  string
		LastDateToApply  string
		Responsibilities string
		LinkBySociety    string
	}

	var role []roles
	if err := database.DB.Model(&models.SocietyRole{}).Select("role_id, role_type, rolename, role_description, last_date_to_apply, responsibilities, link_by_society").Where("society_id = ?", societyID).Scan(&role).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var role models.SocietyRole
	if result := database.DB.First(&role, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}

func RemoveRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	roleID := vars["roleID"]

	if err := database.DB.Where("role_id = ?", roleID).Delete(&models.SocietyRole{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Role successfully deleted"})
}
