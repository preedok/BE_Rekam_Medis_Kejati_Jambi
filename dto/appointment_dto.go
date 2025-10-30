package dto

type CreateAppointmentRequest struct {
	PatientID string `json:"patient_id" validate:"required"`
	DoctorID  string `json:"doctor_id" validate:"required"`
	Poli      string `json:"poli" validate:"required"`
	Date      string `json:"date" validate:"required"`
	Time      string `json:"time" validate:"required"`
	Complaint string `json:"complaint" validate:"required"`
}

type UpdateAppointmentStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=approved rejected"`
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
	})
}
