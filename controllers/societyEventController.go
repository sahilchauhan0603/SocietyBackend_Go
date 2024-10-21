package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	"github.com/sahilchauhan0603/society/models"
)

type tempoAdminEvents struct {
	EventID       uint
	Title         string
	Description   string
	EventDateTime time.Time
	SocietyName   string
}

type event struct {
	SocietyName string
	SocietyID   uint
	EventID     uint
	Title       string
	Description string
	EventType   string
	ModeOfEvent string
	Location    string
	LinkToEvent string
	EventDateTime string
}

func AddNewEvent(w http.ResponseWriter, r *http.Request) {
	var event models.SocietyEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyEvent{}) {
		if err := database.DB.AutoMigrate(&models.SocietyEvent{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&event).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
func UpdateEvent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["eventID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var event models.SocietyEvent
	if result := database.DB.First(&event, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func FetchAllEvents(w http.ResponseWriter, r *http.Request) {
	// var events []models.SocietyEvent
	// if err := database.DB.Order("event_id ASC").Find(&events).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(events)

	var info []event
	if err := database.DB.Table("society_events").
		Select("society_profiles.society_name, society_events.society_id, society_events.event_id, society_events.title, society_events.description, society_events.event_type, society_events.mode_of_event, society_events.location, society_events.link_to_event, society_events.event_date_time").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_events.society_id").
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func FetchEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, err := strconv.Atoi(vars["eventID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var event []models.SocietyEvent
	if err := database.DB.Where("event_id = ?", eventID).First(&event).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func FetchEventsBySocietyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID := vars["societyID"]
	
	var info []event
	if err := database.DB.Table("society_events").
		Select("society_profiles.society_name, society_events.society_id, society_events.event_id, society_events.title, society_events.description, society_events.event_type, society_events.mode_of_event, society_events.location, society_events.link_to_event, society_events.event_date_time").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_events.society_id").Where("society_events.society_id = ?", societyID).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
func RemoveEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, err := strconv.Atoi(vars["eventID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("event_id = ?", eventID).Delete(&models.SocietyEvent{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Event successfully deleted"})
}
func RemoveEventsBySocietyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["societyID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.SocietyEvent{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Events successfully deleted"})
}

// ADMIN PANEL
func FetchAllAdminEvents(w http.ResponseWriter, r *http.Request) {

	var data []tempoAdminEvents
	if err := database.DB.Table("society_events").
		Select("society_events.event_id, society_events.title, society_events.description, society_events.event_date_time, society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_events.society_id").
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func FetchAllAdminEventsSociety(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	var data []tempoAdminEvents
	if err := database.DB.Table("society_events").
		Select("society_events.event_id, society_events.title, society_events.description, society_events.event_date_time, society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_events.society_id").
		Where("society_events.society_id = ?", societyID).
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
