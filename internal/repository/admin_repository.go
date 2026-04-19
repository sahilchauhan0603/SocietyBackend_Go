package repository

import (
	"github.com/sahilchauhan0603/society/internal/database"
	"github.com/sahilchauhan0603/society/internal/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetByRole(role string) (*models.AdminPanelRole, error)
	GetByUsername(username string) (*models.AdminPanelRole, error)
	List() ([]models.AdminPanelRole, error)
	Create(role *models.AdminPanelRole) error
	Update(role *models.AdminPanelRole) error
	DeleteByUsername(username string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{db: database.DB}
}

func (r *adminRepository) GetByRole(role string) (*models.AdminPanelRole, error) {
	var out models.AdminPanelRole
	if err := r.db.Where("role = ?", role).First(&out).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *adminRepository) GetByUsername(username string) (*models.AdminPanelRole, error) {
	var out models.AdminPanelRole
	if err := r.db.Where("username = ?", username).First(&out).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *adminRepository) List() ([]models.AdminPanelRole, error) {
	var out []models.AdminPanelRole
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *adminRepository) Create(role *models.AdminPanelRole) error {
	return r.db.Create(role).Error
}

func (r *adminRepository) Update(role *models.AdminPanelRole) error {
	return r.db.Save(role).Error
}

func (r *adminRepository) DeleteByUsername(username string) error {
	return r.db.Delete(&models.AdminPanelRole{}, "username = ?", username).Error
}
