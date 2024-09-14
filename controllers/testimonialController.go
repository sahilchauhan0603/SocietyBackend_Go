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
	// TestimonialID          uint
	FirstName              string
	LastName               string
	EnrollmentNo           uint
	Branch                 string
	BatchYear              int
	ProfilePicture         string
	SocietyID              uint
	SocietyPosition        string
	TestimonialDescription string
}

type tempAdmin struct {
	TestimonialID          uint
	SocietyID              uint
	EnrollmentNo           uint
	TestimonialDescription string
}

func AddNewTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial models.SocietyTestimonial
	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyTestimonial{}) {
		if err := database.DB.AutoMigrate(&models.SocietyTestimonial{}); err != nil {
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
	id, err := strconv.Atoi(params["testimonialID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var testimonial models.SocietyTestimonial
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
	if err := database.DB.Table("society_testimonials").
		Select("student_profiles.first_name, student_profiles.last_name, society_testimonials.enrollment_no, student_profiles.branch, student_profiles.batch_year, student_profiles.profile_picture, student_profiles.society_id, student_profiles.society_position,society_testimonials.testimonial_description").
		Joins("JOIN student_profiles ON student_profiles.enrollment_no = society_testimonials.enrollment_no").
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

	var info []temp
	if err := database.DB.Table("society_testimonials").
		Select("student_profiles.first_name, student_profiles.last_name, society_testimonials.enrollment_no, society_testimonials.testimonial_id, student_profiles.branch, student_profiles.batch_year, student_profiles.profile_picture, student_profiles.society_id, student_profiles.society_position, society_testimonials.testimonial_description").
		Joins("JOIN student_profiles ON student_profiles.enrollment_no = society_testimonials.enrollment_no").Where("society_testimonials.enrollment_no = ?", enrollmentNo).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func FetchTestimonialBySocietyID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	var info []temp
	if err := database.DB.Table("society_testimonials").
		Select("student_profiles.first_name, student_profiles.last_name, society_testimonials.society_id, student_profiles.branch, student_profiles.batch_year, student_profiles.profile_picture, student_profiles.society_id, student_profiles.society_position, society_testimonials.testimonial_description").
		Joins("JOIN student_profiles ON student_profiles.enrollment_no = society_testimonials.enrollment_no").Where("society_testimonials.society_id = ?", societyID).
		Scan(&info).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func RemoveTestimonialSocietyID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	if err := database.DB.Where("society_id = ?", societyID).Delete(&models.SocietyTestimonial{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Testimonial successfully deleted"})
}

func RemoveTestimonial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testimonialID := vars["testimonialID"]

	if err := database.DB.Where("testimonial_id = ?", testimonialID).Delete(&models.SocietyTestimonial{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Testimonial successfully deleted"})
}



// ADMIN PANEL
func FetchAllTestimonialsAdmin(w http.ResponseWriter, r *http.Request) {

	var data []tempAdmin
	if err := database.DB.Table("society_testimonials").
		Select("society_testimonials.society_id, student_profiles.enrollment_no, society_testimonials.testimonial_description, society_testimonials.testimonial_id").
		Joins("JOIN student_profiles ON student_profiles.enrollment_no = society_testimonials.enrollment_no").
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func FetchAllTestimonialsSocietyAdmin(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	societyID := vars["societyID"]

	var data []tempAdmin
	if err := database.DB.Table("society_testimonials").
		Select("society_testimonials.society_id, student_profiles.enrollment_no, society_testimonials.testimonial_description, society_testimonials.testimonial_id").
		Joins("JOIN student_profiles ON student_profiles.enrollment_no = society_testimonials.enrollment_no").
		Where("society_testimonials.society_id = ?", societyID).
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
