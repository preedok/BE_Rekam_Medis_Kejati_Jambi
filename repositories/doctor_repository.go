package repositories

import (
	"medical-record-api/models"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	FindByEmail(email string) (*models.Doctor, error)
	FindByID(id string) (*models.Doctor, error)
	FindAll() ([]models.Doctor, error)
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{db}
}

func (r *doctorRepository) FindByEmail(email string) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Where("email = ?", email).First(&doctor).Error
	return &doctor, err
}

func (r *doctorRepository) FindByID(id string) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.First(&doctor, "id = ?", id).Error
	return &doctor, err
}

func (r *doctorRepository) FindAll() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := r.db.Find(&doctors).Error
	return doctors, err
}
