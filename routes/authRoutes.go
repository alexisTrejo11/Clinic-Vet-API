package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authController controller.AuthClientController) {
	authClientV1 := app.Group("/api/v1/auth")
	authClientV1.Post("/signUp", authController.ClientSignUp())
	authClientV1.Post("/login", authController.ClientLogin())
}
