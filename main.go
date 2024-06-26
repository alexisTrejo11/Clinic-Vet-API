package main

import (
	"log"

	"example.com/at/backend/api-vet/controller"
	"example.com/at/backend/api-vet/db"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/routes"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Home!")
	})

	dbConn := db.InitDb()

	// Owner
	queries := sqlc.New(dbConn)
	ownerRepository := repository.NewOwnerRepository(queries)
	ownerServices := services.NewOwnerRepository(ownerRepository)
	ownerController := controller.NewOwnerRepository(ownerServices)

	// Routes
	routes.OwnerRoutes(app, ownerController)

	port := ":8080"

	log.Fatal(app.Listen(port))
}
