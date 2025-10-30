package handlers

import (
	"medical-record-api/dto"
	"medical-record-api/services"
	"medical-record-api/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	patientService services.PatientService
	doctorService  services.DoctorService
}

func NewAuthHandler(patientService services.PatientService, doctorService services.DoctorService) *AuthHandler {
	return &AuthHandler{patientService, doctorService}
}

// RegisterPatient godoc
// @Summary Register new patient
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterPatientRequest true "Register Request"
// @Success 201 {object} utils.Response
// @Router /api/auth/register [post]
func (h *AuthHandler) RegisterPatient(c *fiber.Ctx) error {
	var req dto.RegisterPatientRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	patient, err := h.patientService.Register(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "Patient registered successfully", patient)
}

// LoginPatient godoc
// @Summary Patient login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} utils.Response
// @Router /api/auth/login/patient [post]
func (h *AuthHandler) LoginPatient(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.patientService.Login(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.SuccessResponse(c, "Login successful", result)
}

// LoginDoctor godoc
// @Summary Doctor login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} utils.Response
// @Router /api/auth/login/doctor [post]
func (h *AuthHandler) LoginDoctor(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.doctorService.Login(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.SuccessResponse(c, "Login successful", result)
}
