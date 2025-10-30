package services

import (
	"errors"
	"medical-record-api/dto"
	"medical-record-api/models"
	"medical-record-api/repositories"
	"medical-record-api/utils"
	"time"

	"github.com/google/uuid"
)

type PatientService interface {
	Register(req *dto.RegisterPatientRequest) (*models.Patient, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetAll() ([]models.Patient, error)
	GetByID(id string) (*models.Patient, error)
}

type patientService struct {
	repo repositories.PatientRepository
}

func NewPatientService(repo repositories.PatientRepository) PatientService {
	return &patientService{repo}
}

func (s *patientService) Register(req *dto.RegisterPatientRequest) (*models.Patient, error) {
	// Check if email already exists
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	patient := &models.Patient{
		ID:             uuid.New().String(),
		Name:           req.Name,
		Email:          req.Email,
		Password:       hashedPassword,
		Phone:          req.Phone,
		Address:        req.Address,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		BloodType:      req.BloodType,
		RegisteredDate: time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = s.repo.Create(patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *patientService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	patient, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, patient.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(patient.ID, patient.Email, "patient")
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User:  patient,
	}, nil
}

func (s *patientService) GetAll() ([]models.Patient, error) {
	return s.repo.FindAll()
}

func (s *patientService) GetByID(id string) (*models.Patient, error) {
	return s.repo.FindByID(id)
}
