package repositories

import (
	"medical-record-api/models"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
	FindByID(id string) (*models.Appointment, error)
	FindAll() ([]models.Appointment, error)
	FindByPatientID(patientID string) ([]models.Appointment, error)
	FindByDoctorID(doctorID string) ([]models.Appointment, error)
	Update(appointment *models.Appointment) error
	Delete(id string) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db}
}

func (r *appointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) FindByID(id string) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.First(&appointment, "id = ?", id).Error
	return &appointment, err
}

func (r *appointmentRepository) FindAll() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Order("date DESC, time DESC").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) FindByPatientID(patientID string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Where("patient_id = ?", patientID).Order("date DESC").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) FindByDoctorID(doctorID string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Where("doctor_id = ?", doctorID).Order("date DESC").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) Update(appointment *models.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *appointmentRepository) Delete(id string) error {
	return r.db.Delete(&models.Appointment{}, "id = ?", id).Error
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
