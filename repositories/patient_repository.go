package repositories

import (
	"medical-record-api/models"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	FindByEmail(email string) (*models.Patient, error)
	FindByID(id string) (*models.Patient, error)
	FindAll() ([]models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id string) error
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) FindByEmail(email string) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Where("email = ?", email).First(&patient).Error
	return &patient, err
}

func (r *patientRepository) FindByID(id string) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, "id = ?", id).Error
	return &patient, err
}

func (r *patientRepository) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Find(&patients).Error
	return patients, err
}

func (r *patientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id string) error {
	return r.db.Delete(&models.Patient{}, "id = ?", id).Error
}

func DoctorOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role != "admin" {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Access denied: doctors only")
		}
		return c.Next()
	}
}
