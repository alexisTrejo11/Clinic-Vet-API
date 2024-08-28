package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authClientController *controller.AuthClientController, authEmployeeController *controller.AuthEmployeeController) {
	authClientV1 := app.Group("v1/api/clinic-vet")
	authClientV1.Post("/signup", authClientController.ClientSignUp())
	authClientV1.Post("/login", authClientController.ClientLogin())

	authEmployeeV1 := app.Group("/v1/api/employee/clinic-vet")
	authEmployeeV1.Post("/signup", authEmployeeController.EmployeeSignUp())
	authEmployeeV1.Post("/login", authEmployeeController.EmployeeLogin())
}
