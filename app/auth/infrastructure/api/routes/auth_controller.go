package routes

import (
	"example.com/at/backend/api-vet/app/accounts/infrastructure/api/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authClientController *controller.AuthClientController, authEmployeeController *controller.AuthEmployeeController) {
	authClientV2 := app.Group("v2/api/clinic-vet")
	authClientV2.Post("/signup", authClientController.ClientSignUp())
	authClientV2.Post("/login", authClientController.ClientLogin())

	authEmployeeV2 := app.Group("/v2/api/employee/clinic-vet")
	authEmployeeV2.Post("/signup", authEmployeeController.EmployeeSignUp())
	authEmployeeV2.Post("/login", authEmployeeController.EmployeeLogin())
}
s