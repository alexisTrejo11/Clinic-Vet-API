package main

import (
	"log"

	"example.com/at/backend/api-vet/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Home!")
	})

	port := ":8080"

	log.Fatal(app.Listen(port))
}
