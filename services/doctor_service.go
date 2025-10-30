package services

import (
	"errors"
	"medical-record-api/dto"
	"medical-record-api/models"
	"medical-record-api/repositories"
	"medical-record-api/utils"
)

type DoctorService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetAll() ([]models.Doctor, error)
	GetByID(id string) (*models.Doctor, error)
}

type doctorService struct {
	repo repositories.DoctorRepository
}

func NewDoctorService(repo repositories.DoctorRepository) DoctorService {
	return &doctorService{repo}
}

func (s *doctorService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	doctor, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, doctor.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(doctor.ID, doctor.Email, "admin")
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User:  doctor,
	}, nil
}

func (s *doctorService) GetAll() ([]models.Doctor, error) {
	return s.repo.FindAll()
}

func (s *doctorService) GetByID(id string) (*models.Doctor, error) {
	return s.repo.FindByID(id)
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
