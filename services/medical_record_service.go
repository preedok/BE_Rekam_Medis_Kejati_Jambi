package services

import (
	"medical-record-api/dto"
	"medical-record-api/models"
	"medical-record-api/repositories"
	"time"

	"github.com/google/uuid"
)

type MedicalRecordService interface {
	Create(req *dto.CreateMedicalRecordRequest, createdBy string) (*models.MedicalRecord, error)
	GetAll() ([]models.MedicalRecord, error)
	GetByID(id string) (*models.MedicalRecord, error)
	GetByPatientID(patientID string) ([]models.MedicalRecord, error)
	GetByCreatedBy(createdBy string) ([]models.MedicalRecord, error)
	GetByMonth(month string) ([]models.MedicalRecord, error)
	Delete(id string) error
}

type medicalRecordService struct {
	repo            repositories.MedicalRecordRepository
	appointmentRepo repositories.AppointmentRepository
	patientRepo     repositories.PatientRepository
}

func NewMedicalRecordService(
	repo repositories.MedicalRecordRepository,
	appointmentRepo repositories.AppointmentRepository,
	patientRepo repositories.PatientRepository,
) MedicalRecordService {
	return &medicalRecordService{repo, appointmentRepo, patientRepo}
}

func (s *medicalRecordService) Create(req *dto.CreateMedicalRecordRequest, createdBy string) (*models.MedicalRecord, error) {
	// Get appointment info
	appointment, err := s.appointmentRepo.FindByID(req.AppointmentID)
	if err != nil {
		return nil, err
	}

	// Get patient info
	patient, err := s.patientRepo.FindByID(req.PatientID)
	if err != nil {
		return nil, err
	}

	record := &models.MedicalRecord{
		ID:            uuid.New().String(),
		AppointmentID: req.AppointmentID,
		PatientID:     req.PatientID,
		PatientName:   patient.Name,
		Poli:          appointment.Poli,
		Date:          appointment.Date,
		Anamnesa:      req.Anamnesa,
		Objective:     req.Objective,
		Diagnosis:     req.Diagnosis,
		Therapy:       req.Therapy,
		Prescription:  req.Prescription,
		NextVisit:     req.NextVisit,
		CreatedBy:     createdBy,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = s.repo.Create(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s *medicalRecordService) GetAll() ([]models.MedicalRecord, error) {
	return s.repo.FindAll()
}

func (s *medicalRecordService) GetByID(id string) (*models.MedicalRecord, error) {
	return s.repo.FindByID(id)
}

func (s *medicalRecordService) GetByPatientID(patientID string) ([]models.MedicalRecord, error) {
	return s.repo.FindByPatientID(patientID)
}

func (s *medicalRecordService) GetByCreatedBy(createdBy string) ([]models.MedicalRecord, error) {
	return s.repo.FindByCreatedBy(createdBy)
}

func (s *medicalRecordService) GetByMonth(month string) ([]models.MedicalRecord, error) {
	return s.repo.FindByMonth(month)
}

func (s *medicalRecordService) Delete(id string) error {
	return s.repo.Delete(id)
}
