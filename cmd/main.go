package main

import (
	"log"
	"medical-record-api/config"
	"medical-record-api/handlers"
	"medical-record-api/middleware"
	"medical-record-api/repositories"
	"medical-record-api/routes"
	"medical-record-api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Medical Record API
// @version 1.0
// @description API for Medical Record Management System
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to database
	config.ConnectDatabase()

	// Initialize repositories
	patientRepo := repositories.NewPatientRepository(config.DB)
	doctorRepo := repositories.NewDoctorRepository(config.DB)
	appointmentRepo := repositories.NewAppointmentRepository(config.DB)
	medicalRecordRepo := repositories.NewMedicalRecordRepository(config.DB)

	// Initialize services
	patientService := services.NewPatientService(patientRepo)
	doctorService := services.NewDoctorService(doctorRepo)
	appointmentService := services.NewAppointmentService(appointmentRepo, patientRepo, doctorRepo)
	medicalRecordService := services.NewMedicalRecordService(medicalRecordRepo, appointmentRepo, patientRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(patientService, doctorService)
	patientHandler := handlers.NewPatientHandler(patientService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	medicalRecordHandler := handlers.NewMedicalRecordHandler(medicalRecordService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(middleware.CORS())

	// Setup routes
	routes.SetupRoutes(app, authHandler, patientHandler, appointmentHandler, medicalRecordHandler)

	// Start server
	port := ":" + config.AppConfig.AppPort
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(port))
}
