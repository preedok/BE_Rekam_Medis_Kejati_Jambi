package handlers

import (
	"medical-record-api/dto"
	"medical-record-api/services"
	"medical-record-api/utils"

	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	service services.AppointmentService
}

func NewAppointmentHandler(service services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service}
}

// CreateAppointment godoc
// @Summary Create appointment
// @Tags Appointments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateAppointmentRequest true "Appointment Request"
// @Success 201 {object} utils.Response
// @Router /api/appointments [post]
func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	appointment, err := h.service.Create(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "Appointment created successfully", appointment)
}

// GetAllAppointments godoc
// @Summary Get all appointments
// @Tags Appointments
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/appointments [get]
func (h *AppointmentHandler) GetAll(c *fiber.Ctx) error {
	appointments, err := h.service.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Appointments retrieved successfully", appointments)
}

// GetAppointmentsByDoctor godoc
// @Summary Get appointments by doctor
// @Tags Appointments
// @Security BearerAuth
// @Produce json
// @Param doctorId path string true "Doctor ID"
// @Success 200 {object} utils.Response
// @Router /api/appointments/doctor/{doctorId} [get]
func (h *AppointmentHandler) GetByDoctorID(c *fiber.Ctx) error {
	doctorID := c.Params("doctorId")
	appointments, err := h.service.GetByDoctorID(doctorID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Appointments retrieved successfully", appointments)
}

// UpdateAppointmentStatus godoc
// @Summary Update appointment status
// @Tags Appointments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Appointment ID"
// @Param request body dto.UpdateAppointmentStatusRequest true "Status Request"
// @Success 200 {object} utils.Response
// @Router /api/appointments/{id}/status [patch]
func (h *AppointmentHandler) UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateAppointmentStatusRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	appointment, err := h.service.UpdateStatus(id, req.Status)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "Appointment status updated successfully", appointment)
}
