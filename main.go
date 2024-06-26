package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Home!")
	})

	port := ":8080"

	log.Fatal(app.Listen(port))
}
