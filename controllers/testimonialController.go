package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	"github.com/sahilchauhan0603/society/models"
)

type temp struct {
	TestimonialDescription string
	FirstName              string
	LastName               string
	Branch                 string
	BatchYear              int
	ProfilePicture         string
	SocietyID              uint
	SocietyPosition        string
}

func AddNewTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial models.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Testimonial{}) {
		if err := database.DB.AutoMigrate(&models.Testimonial{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&testimonial).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testimonial)
}
func UpdateTestimonial(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["enrollmentNo"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var testimonial models.Testimonial
	if result := database.DB.First(&testimonial, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&testimonial)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(testimonial)
}
func FetchAllTestimonials(w http.ResponseWriter, r *http.Request) {
	var data []temp
	if err := database.DB.Table("testimonials").
		Select("testimonials.testimonial_description, student_profiles.first_name, student_profiles.last_name, student_profiles.branch, student_profiles.batch_year, student_profiles.profile_picture, student_profiles.society_id, student_profiles.society_position").
		Joins("JOIN student_profiles ON student_profiles.enrollment_no = testimonials.enrollment_no").
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func FetchTestimonialByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo := vars["enrollmentNo"]

	var info temp
	if err := database.DB.Table("testimonials").
	Select("testimonials.testimonial_description, student_profiles.first_name, student_profiles.last_name, student_profiles.branch, student_profiles.batch_year, student_profiles.profile_picture, student_profiles.society_id, student_profiles.society_position").
	Joins("JOIN student_profiles ON student_profiles.enrollment_no = testimonials.enrollment_no").Where("testimonials.enrollment_no = ?", enrollmentNo).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func RemoveTestimonial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enrollmentNo := vars["enrollmentNo"]

	if err := database.DB.Where("enrollment_no = ?", enrollmentNo).Delete(&models.Testimonial{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Testimonial successfully deleted"})
}
