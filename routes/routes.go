package routes

import (
	"medical-record-api/handlers"
	"medical-record-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(
	app *fiber.App,
	authHandler *handlers.AuthHandler,
	patientHandler *handlers.PatientHandler,
	appointmentHandler *handlers.AppointmentHandler,
	medicalRecordHandler *handlers.MedicalRecordHandler,
) {
	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Medical Record API is running",
			"version": "1.0.0",
		})
	})

	// API routes
	api := app.Group("/api")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.RegisterPatient)
	auth.Post("/login/patient", authHandler.LoginPatient)
	auth.Post("/login/doctor", authHandler.LoginDoctor)

	// Protected routes
	protected := api.Group("", middleware.AuthRequired())

	// Patient routes
	patients := protected.Group("/patients")
	patients.Get("/", patientHandler.GetAll)
	patients.Get("/:id", patientHandler.GetByID)

	// Appointment routes
	appointments := protected.Group("/appointments")
	appointments.Post("/", appointmentHandler.Create)
	appointments.Get("/", appointmentHandler.GetAll)
	appointments.Get("/doctor/:doctorId", appointmentHandler.GetByDoctorID)
	appointments.Patch("/:id/status", appointmentHandler.UpdateStatus)

	// Medical Record routes
	medicalRecords := protected.Group("/medical-records")
	medicalRecords.Post("/", middleware.DoctorOnly(), medicalRecordHandler.Create)
	medicalRecords.Get("/", medicalRecordHandler.GetAll)
	medicalRecords.Get("/patient/:patientId", medicalRecordHandler.GetByPatientID)
	medicalRecords.Get("/report/monthly", middleware.DoctorOnly(), medicalRecordHandler.GetMonthlyReport)
}
