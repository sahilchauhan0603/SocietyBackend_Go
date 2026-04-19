package repository

import (
	"github.com/sahilchauhan0603/society/internal/database"
	"github.com/sahilchauhan0603/society/internal/models"
	"gorm.io/gorm"
)

type ContentRepository interface {
	ListNews() ([]models.SocietyNews, error)
	ListGalleries() ([]models.SocietyGallery, error)
	ListTestimonials() ([]models.SocietyTestimonial, error)
	CreateNews(news *models.SocietyNews) error
	CreateGallery(gallery *models.SocietyGallery) error
	CreateTestimonial(testimonial *models.SocietyTestimonial) error
}

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository() ContentRepository {
	return &contentRepository{db: database.DB}
}

func (r *contentRepository) ListNews() ([]models.SocietyNews, error) {
	var out []models.SocietyNews
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *contentRepository) ListGalleries() ([]models.SocietyGallery, error) {
	var out []models.SocietyGallery
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *contentRepository) ListTestimonials() ([]models.SocietyTestimonial, error) {
	var out []models.SocietyTestimonial
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *contentRepository) CreateNews(news *models.SocietyNews) error {
	return r.db.Create(news).Error
}

func (r *contentRepository) CreateGallery(gallery *models.SocietyGallery) error {
	return r.db.Create(gallery).Error
}

func (r *contentRepository) CreateTestimonial(testimonial *models.SocietyTestimonial) error {
	return r.db.Create(testimonial).Error
}
