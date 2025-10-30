package dto

type CreateMedicalRecordRequest struct {
	AppointmentID string `json:"appointment_id" validate:"required"`
	PatientID     string `json:"patient_id" validate:"required"`
	Anamnesa      string `json:"anamnesa" validate:"required"`
	Objective     string `json:"objective" validate:"required"`
	Diagnosis     string `json:"diagnosis" validate:"required"`
	Therapy       string `json:"therapy" validate:"required"`
	Prescription  string `json:"prescription" validate:"required"`
	NextVisit     string `json:"next_visit"`
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
	})
}
