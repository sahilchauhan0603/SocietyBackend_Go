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

type tempNews struct {
	SocietyName string
	SocietyID   uint
	NewsID      uint
	Title       string
	Description string
	DateOfNews  time.Time
	Author      string
}
type tempAdminNews struct {
	DateOfNews time.Time
	Author     string
	Title      string
}

type tempNewsSocietyAdmin struct {
	SocietyID   uint
	NewsID      uint
	Title       string
	DateOfNews  time.Time
	SocietyName string
	Author      string
	Description string
}

func AddNewNews(w http.ResponseWriter, r *http.Request) {

	var news models.SocietyNews
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyNews{}) {
		if err := database.DB.AutoMigrate(&models.SocietyNews{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := database.DB.Create(&news).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(news)
}

func FetchAllNews(w http.ResponseWriter, r *http.Request) {

	var newsList []tempNews
	if result := database.DB.Table("society_news").
		Select("society_news.*,society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_news.society_id").
		Scan(&newsList); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newsList)
}

func FetchNews(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["society_id"])
	if err != nil {
		http.Error(w, "Invalid Society ID", http.StatusBadRequest)
		return
	}

	var info []tempNews
	if err := database.DB.Table("society_news").
		Select("society_news.*,society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_news.society_id").Where("society_news.society_id = ?", societyID).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["newsID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var news models.SocietyNews
	if result := database.DB.First(&news, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&news)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}

func RemoveNews(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	newsID, err := strconv.Atoi(vars["newsID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Where("news_id = ?", newsID).Delete(&models.SocietyNews{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "News successfully deleted"})
}

// ADMIN PANEL
func FetchAllNewsAdminHome(w http.ResponseWriter, r *http.Request) {

	var newsList []tempAdminNews
	if result := database.DB.Table("society_news").
		Select("society_news.date_of_news, society_news.author, society_news.title").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_news.society_id").
		Scan(&newsList); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newsList)
}

func FetchAllNewsAdminNews(w http.ResponseWriter, r *http.Request) {

	var info []tempNewsSocietyAdmin
	if err := database.DB.Table("society_news").
		Select("society_news.society_id, society_news.news_id, society_news.title, society_news.date_of_news, society_news.description, society_news.author ,society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_news.society_id").
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}

func FetchNewsAdminNews(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID, err := strconv.Atoi(vars["society_id"])
	if err != nil {
		http.Error(w, "Invalid Society ID", http.StatusBadRequest)
		return
	}

	var info []tempNewsSocietyAdmin
	if err := database.DB.Table("society_news").
		Select("society_news.society_id, society_news.news_id, society_news.title, society_news.date_of_news, society_news.description, society_news.author ,society_profiles.society_name").
		Joins("JOIN society_profiles ON society_profiles.society_id = society_news.society_id").Where("society_news.society_id = ?", societyID).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}
