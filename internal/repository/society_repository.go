package repository

import (
	"github.com/sahilchauhan0603/society/internal/database"
	"github.com/sahilchauhan0603/society/internal/models"
	"gorm.io/gorm"
)

type SocietyRepository interface {
	ListSocieties() ([]models.SocietyProfile, error)
	GetSocietyByID(id uint) (*models.SocietyProfile, error)
	CreateSociety(society *models.SocietyProfile) error
	UpdateSociety(society *models.SocietyProfile) error
	DeleteSocietyByID(id uint) error
	ListRoles() ([]models.SocietyRole, error)
	ListEvents() ([]models.SocietyEvent, error)
	ListAchievements() ([]models.SocietyAchievement, error)
}

type societyRepository struct {
	db *gorm.DB
}

func NewSocietyRepository() SocietyRepository {
	return &societyRepository{db: database.DB}
}

func (r *societyRepository) ListSocieties() ([]models.SocietyProfile, error) {
	var out []models.SocietyProfile
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *societyRepository) GetSocietyByID(id uint) (*models.SocietyProfile, error) {
	var out models.SocietyProfile
	if err := r.db.First(&out, "society_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *societyRepository) CreateSociety(society *models.SocietyProfile) error {
	return r.db.Create(society).Error
}

func (r *societyRepository) UpdateSociety(society *models.SocietyProfile) error {
	return r.db.Save(society).Error
}

func (r *societyRepository) DeleteSocietyByID(id uint) error {
	return r.db.Delete(&models.SocietyProfile{}, "society_id = ?", id).Error
}

func (r *societyRepository) ListRoles() ([]models.SocietyRole, error) {
	var out []models.SocietyRole
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *societyRepository) ListEvents() ([]models.SocietyEvent, error) {
	var out []models.SocietyEvent
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *societyRepository) ListAchievements() ([]models.SocietyAchievement, error) {
	var out []models.SocietyAchievement
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}
