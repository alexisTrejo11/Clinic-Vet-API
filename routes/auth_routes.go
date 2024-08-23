package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authClientController *controller.AuthClientController, authEmployeeController *controller.AuthEmployeeController) {
	authClientV1 := app.Group("/api/v1/auth")
	authClientV1.Post("/signUp", authClientController.ClientSignUp())
	authClientV1.Post("/login", authClientController.ClientLogin())

	authEmployeeV1 := app.Group("/api/v1/employee/auth")
	authEmployeeV1.Post("/signUp", authEmployeeController.EmployeeSignUp())
	authEmployeeV1.Post("/login", authEmployeeController.EmployeeLogin())
}
