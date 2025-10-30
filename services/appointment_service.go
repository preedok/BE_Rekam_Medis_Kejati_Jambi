package services

import (
	"medical-record-api/dto"
	"medical-record-api/models"
	"medical-record-api/repositories"
	"time"

	"github.com/google/uuid"
)

type AppointmentService interface {
	Create(req *dto.CreateAppointmentRequest) (*models.Appointment, error)
	GetAll() ([]models.Appointment, error)
	GetByID(id string) (*models.Appointment, error)
	GetByPatientID(patientID string) ([]models.Appointment, error)
	GetByDoctorID(doctorID string) ([]models.Appointment, error)
	UpdateStatus(id string, status string) (*models.Appointment, error)
	Delete(id string) error
}

type appointmentService struct {
	repo        repositories.AppointmentRepository
	patientRepo repositories.PatientRepository
	doctorRepo  repositories.DoctorRepository
}

func NewAppointmentService(
	repo repositories.AppointmentRepository,
	patientRepo repositories.PatientRepository,
	doctorRepo repositories.DoctorRepository,
) AppointmentService {
	return &appointmentService{repo, patientRepo, doctorRepo}
}

func (s *appointmentService) Create(req *dto.CreateAppointmentRequest) (*models.Appointment, error) {
	// Get patient and doctor info
	patient, err := s.patientRepo.FindByID(req.PatientID)
	if err != nil {
		return nil, err
	}

	doctor, err := s.doctorRepo.FindByID(req.DoctorID)
	if err != nil {
		return nil, err
	}

	appointment := &models.Appointment{
		ID:          uuid.New().String(),
		PatientID:   req.PatientID,
		PatientName: patient.Name,
		DoctorID:    req.DoctorID,
		DoctorName:  doctor.Name,
		Poli:        req.Poli,
		Date:        req.Date,
		Time:        req.Time,
		Complaint:   req.Complaint,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = s.repo.Create(appointment)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (s *appointmentService) GetAll() ([]models.Appointment, error) {
	return s.repo.FindAll()
}

func (s *appointmentService) GetByID(id string) (*models.Appointment, error) {
	return s.repo.FindByID(id)
}

func (s *appointmentService) GetByPatientID(patientID string) ([]models.Appointment, error) {
	return s.repo.FindByPatientID(patientID)
}

func (s *appointmentService) GetByDoctorID(doctorID string) ([]models.Appointment, error) {
	return s.repo.FindByDoctorID(doctorID)
}

func (s *appointmentService) UpdateStatus(id string, status string) (*models.Appointment, error) {
	appointment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	appointment.Status = status
	appointment.UpdatedAt = time.Now()

	err = s.repo.Update(appointment)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (s *appointmentService) Delete(id string) error {
	return s.repo.Delete(id)
}
