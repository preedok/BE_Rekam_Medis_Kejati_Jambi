package handlers

import (
	"medical-record-api/dto"
	"medical-record-api/services"
	"medical-record-api/utils"

	"github.com/gofiber/fiber/v2"
)

type MedicalRecordHandler struct {
	service services.MedicalRecordService
}

func NewMedicalRecordHandler(service services.MedicalRecordService) *MedicalRecordHandler {
	return &MedicalRecordHandler{service}
}

// CreateMedicalRecord godoc
// @Summary Create medical record
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateMedicalRecordRequest true "Medical Record Request"
// @Success 201 {object} utils.Response
// @Router /api/medical-records [post]
func (h *MedicalRecordHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateMedicalRecordRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	createdBy := c.Locals("email").(string)
	record, err := h.service.Create(&req, createdBy)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, "Medical record created successfully", record)
}

// GetAllMedicalRecords godoc
// @Summary Get all medical records
// @Tags Medical Records
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/medical-records [get]
func (h *MedicalRecordHandler) GetAll(c *fiber.Ctx) error {
	records, err := h.service.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Medical records retrieved successfully", records)
}

// GetMedicalRecordsByPatient godoc
// @Summary Get medical records by patient
// @Tags Medical Records
// @Security BearerAuth
// @Produce json
// @Param patientId path string true "Patient ID"
// @Success 200 {object} utils.Response
// @Router /api/medical-records/patient/{patientId} [get]
func (h *MedicalRecordHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID := c.Params("patientId")
	records, err := h.service.GetByPatientID(patientID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Medical records retrieved successfully", records)
}

// GetMonthlyReport godoc
// @Summary Get monthly report
// @Tags Medical Records
// @Security BearerAuth
// @Produce json
// @Param month query string true "Month (YYYY-MM)"
// @Success 200 {object} utils.Response
// @Router /api/medical-records/report/monthly [get]
func (h *MedicalRecordHandler) GetMonthlyReport(c *fiber.Ctx) error {
	month := c.Query("month")
	if month == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Month parameter is required")
	}

	records, err := h.service.GetByMonth(month)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, "Monthly report retrieved successfully", records)
}
