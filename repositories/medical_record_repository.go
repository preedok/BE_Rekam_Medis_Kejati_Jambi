package repositories

import (
	"medical-record-api/models"

	"gorm.io/gorm"
)

type MedicalRecordRepository interface {
	Create(record *models.MedicalRecord) error
	FindByID(id string) (*models.MedicalRecord, error)
	FindAll() ([]models.MedicalRecord, error)
	FindByPatientID(patientID string) ([]models.MedicalRecord, error)
	FindByCreatedBy(createdBy string) ([]models.MedicalRecord, error)
	FindByMonth(month string) ([]models.MedicalRecord, error)
	Update(record *models.MedicalRecord) error
	Delete(id string) error
}

type medicalRecordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) MedicalRecordRepository {
	return &medicalRecordRepository{db}
}

func (r *medicalRecordRepository) Create(record *models.MedicalRecord) error {
	return r.db.Create(record).Error
}

func (r *medicalRecordRepository) FindByID(id string) (*models.MedicalRecord, error) {
	var record models.MedicalRecord
	err := r.db.First(&record, "id = ?", id).Error
	return &record, err
}

func (r *medicalRecordRepository) FindAll() ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := r.db.Order("created_at DESC").Find(&records).Error
	return records, err
}

func (r *medicalRecordRepository) FindByPatientID(patientID string) ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := r.db.Where("patient_id = ?", patientID).Order("created_at DESC").Find(&records).Error
	return records, err
}

func (r *medicalRecordRepository) FindByCreatedBy(createdBy string) ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := r.db.Where("created_by = ?", createdBy).Order("created_at DESC").Find(&records).Error
	return records, err
}

func (r *medicalRecordRepository) FindByMonth(month string) ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := r.db.Where("TO_CHAR(created_at, 'YYYY-MM') = ?", month).Find(&records).Error
	return records, err
}

func (r *medicalRecordRepository) Update(record *models.MedicalRecord) error {
	return r.db.Save(record).Error
}

func (r *medicalRecordRepository) Delete(id string) error {
	return r.db.Delete(&models.MedicalRecord{}, "id = ?", id).Error
}
