package handlers

import (
	"medical-record-api/services"
	"medical-record-api/utils"

	"github.com/gofiber/fiber/v2"
)

type PatientHandler struct {
	service services.PatientService
}

func NewPatientHandler(service services.PatientService) *PatientHandler {
	return &PatientHandler{service}
}

// GetAllPatients godoc
// @Summary Get all patients
// @Tags Patients
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/patients [get]
func (h *PatientHandler) GetAll(c *fiber.Ctx) error {
	patients, err := h.service.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Patients retrieved successfully", patients)
}

// GetPatientByID godoc
// @Summary Get patient by ID
// @Tags Patients
// @Security BearerAuth
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {object} utils.Response
// @Router /api/patients/{id} [get]
func (h *PatientHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	patient, err := h.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Patient not found")
	}

	return utils.SuccessResponse(c, "Patient retrieved successfully", patient)
}
