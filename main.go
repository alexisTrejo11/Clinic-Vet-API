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

// @title Vet API
// @version 1.0
// @description This is a sample server for a vet clinic.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	// Server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Home!")
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

	// Routes
	routes.OwnerRoutes(app, ownerController)
	routes.PetsRoutes(app, petController)

	port := ":8000"

	log.Fatal(app.Listen(port))
}
