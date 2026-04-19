package routes

import (
	"github.com/gorilla/mux"
	contentctrl "github.com/sahilchauhan0603/society/internal/http/controllers/content"
)

func registerContentRoutes(api *mux.Router) {
	api.HandleFunc("/testimonials", contentctrl.FetchAllTestimonials).Methods("GET")
	api.HandleFunc("/testimonials/{societyID}", contentctrl.RemoveTestimonialSocietyID).Methods("DELETE")
	api.HandleFunc("/testimonials/{enrollmentNo}", contentctrl.FetchTestimonialByID).Methods("GET")
	api.HandleFunc("/testimonials/society/{societyID}", contentctrl.FetchTestimonialBySocietyID).Methods("GET")

	api.HandleFunc("/galleries", contentctrl.FetchAllGalleries).Methods("GET")
	api.HandleFunc("/galleries/{society_id}", contentctrl.FetchGallerySociety).Methods("GET")

	api.HandleFunc("/news", contentctrl.FetchAllNews).Methods("GET")
	api.HandleFunc("/news/{society_id}", contentctrl.FetchNews).Methods("GET")

	api.HandleFunc("/contact", contentctrl.ContactUSHandler).Methods("POST")
	api.HandleFunc("/feedback", contentctrl.FeedbackHandler).Methods("POST")
	api.HandleFunc("/becomeMember", contentctrl.BecomeMemberHandler).Methods("POST")

	api.HandleFunc("/admin/home/news", contentctrl.FetchAllNewsAdminHome).Methods("GET")
	api.HandleFunc("/news", contentctrl.AddNewNews).Methods("POST")
	api.HandleFunc("/admin/news/{society_id}", contentctrl.FetchNewsAdminNews).Methods("GET")
	api.HandleFunc("/admin/news", contentctrl.FetchAllNewsAdminNews).Methods("GET")
	api.HandleFunc("/news/{newsID}", contentctrl.UpdateNews).Methods("PUT")
	api.HandleFunc("/news/{newsID}", contentctrl.RemoveNews).Methods("DELETE")

	api.HandleFunc("/testimonials", contentctrl.AddNewTestimonial).Methods("POST")
	api.HandleFunc("/testimonials/{testimonialID}", contentctrl.UpdateTestimonial).Methods("PUT")
	api.HandleFunc("/admin/testimonials", contentctrl.FetchAllTestimonialsAdmin).Methods("GET")
	api.HandleFunc("/admin/testimonials/{societyID}", contentctrl.FetchAllTestimonialsSocietyAdmin).Methods("GET")
	api.HandleFunc("/testimonials/{testimonialID}", contentctrl.RemoveTestimonial).Methods("DELETE")

	api.HandleFunc("/galleries", contentctrl.AddNewGallery).Methods("POST")
	api.HandleFunc("/admin/gallery", contentctrl.FetchAllGalleries).Methods("GET")
	api.HandleFunc("/admin/gallery/{society_id}", contentctrl.FetchGallery).Methods("GET")
	api.HandleFunc("/galleries/{societyID}", contentctrl.UpdateGallery).Methods("PUT")
	api.HandleFunc("/galleries/{societyID}", contentctrl.RemoveGallery).Methods("DELETE")
	api.HandleFunc("/galleries/{galleryID}", contentctrl.RemoveGallerySocietyID).Methods("DELETE")
}
