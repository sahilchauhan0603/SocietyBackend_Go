package repository

import (
	"github.com/sahilchauhan0603/society/internal/database"
	"github.com/sahilchauhan0603/society/internal/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	ListStudents() ([]models.StudentProfile, error)
	GetStudentByEnrollment(enrollment uint) (*models.StudentProfile, error)
	CreateStudent(student *models.StudentProfile) error
	UpdateStudent(student *models.StudentProfile) error
	DeleteStudentByEnrollment(enrollment uint) error
	ListStudentAchievements() ([]models.StudentAchievement, error)
	ListStudentMarkings() ([]models.StudentMarking, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{db: database.DB}
}

func (r *studentRepository) ListStudents() ([]models.StudentProfile, error) {
	var out []models.StudentProfile
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *studentRepository) GetStudentByEnrollment(enrollment uint) (*models.StudentProfile, error) {
	var out models.StudentProfile
	if err := r.db.First(&out, "enrollment_no = ?", enrollment).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *studentRepository) CreateStudent(student *models.StudentProfile) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) UpdateStudent(student *models.StudentProfile) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) DeleteStudentByEnrollment(enrollment uint) error {
	return r.db.Delete(&models.StudentProfile{}, "enrollment_no = ?", enrollment).Error
}

func (r *studentRepository) ListStudentAchievements() ([]models.StudentAchievement, error) {
	var out []models.StudentAchievement
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *studentRepository) ListStudentMarkings() ([]models.StudentMarking, error) {
	var out []models.StudentMarking
	if err := r.db.Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}
