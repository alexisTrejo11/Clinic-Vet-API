package main

import (
	"log"

	"example.com/at/backend/api-vet/controller"
	"example.com/at/backend/api-vet/db"
	_ "example.com/at/backend/api-vet/docs"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/routes"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Clinic Vet API
// @version 1.0
// @description This is a sample server for a vet clinic.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.email marcoalexispt.02@gmail.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	// Server
	app := fiber.New()

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("¡Welcome to Vet API!")
	})

	// Serve Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	dbConn := db.InitDb()

	// Database
	queries := sqlc.New(dbConn)

	// Owner
	ownerRepository := repository.NewOwnerRepositoryImpl(queries)
	ownerServices := services.NewOwnerService(ownerRepository)
	ownerController := controller.NewOwnerController(ownerServices)

	// Pet
	petRepository := repository.NewPetRepository(queries)
	petServices := services.NewPetService(petRepository)
	petController := controller.NewPetController(petServices)

	// Appointment
	appointmentRepository := repository.NewAppointmentRepository(queries)
	appointmentService := services.NewAppointmentService(appointmentRepository)
	appointmentController := controller.NewAppointmentController(appointmentService)

	// Owner-Pet
	ownerPetController := controller.NewOwnerPetController(ownerServices, *petServices)

	// Owner-Appointment
	ownerAppController := controller.NewOwnerAppointmentController(appointmentService, ownerServices)

	// Vet
	vetRepository := repository.NewVeterinarianRepository(queries)
	vetServices := services.NewVeterinarianService(vetRepository)
	vetController := controller.NewVeterinarianController(vetServices)

	// Owner-MedicalHistory
	medicalHistoryRepository := repository.NewMedicalHistoryRepository(queries)
	medicalHistoryService := services.NewMedicalHistoryService(medicalHistoryRepository)
	medicalHistoryController := controller.NewClientMedicalHistory(*petServices, ownerServices, medicalHistoryService)

	// Users
	userRepository := repository.NewUserRepository(queries)
	authCommonService := services.NewCommonAuthService(userRepository, ownerRepository, vetRepository)

	// Auth-Client
	authClientService := services.NewClientAuthService(authCommonService, userRepository, ownerRepository, vetRepository)
	authClientController := controller.NewAuthClientController(authClientService, authCommonService)

	// Auth-Employees
	authEmployeeService := services.NewAuthEmployeeService(userRepository, vetRepository, authCommonService)
	authEmployeeController := controller.NewAuthEmployeeController(authEmployeeService, authCommonService, vetServices)

	// Routes
	routes.OwnerRoutes(app, ownerController)
	routes.PetsRoutes(app, petController)
	routes.AppointmentRoutes(app, appointmentController)
	routes.VeterinarianRoutes(app, vetController)
	routes.AuthRoutes(app, authClientController, authEmployeeController)
	routes.OwnerPetRoutes(app, ownerPetController)
	routes.OwnerAppointmentRoutes(app, ownerAppController)
	routes.OwnerMedicalHistoryRoutes(app, medicalHistoryController)

	port := ":8000"

	log.Fatal(app.Listen(port))
}
