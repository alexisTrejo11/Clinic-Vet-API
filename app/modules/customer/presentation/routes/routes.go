// Package routes defines the HTTP routes for customer-related operations.
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/customer/presentation/controller"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(app *gin.RouterGroup, customerController *controller.CustomerController, authMiddleware *middleware.AuthMiddleware) {
	customerGroup := app.Group("/customers")
	//customerGroup.Use(authMiddleware.Authenticate())
	//customerGroup.Use(authMiddleware.RequireAnyRole("admin", "receptionist", "veterinarian"))

	customerGroup.GET("/:id", customerController.GetCustomerByID)
	customerGroup.GET("/", customerController.SearchCustomers)
	customerGroup.POST("/", customerController.CreateCustomer)
	customerGroup.PATCH("/:id", customerController.UpdateCustomer)
	customerGroup.DELETE("/:id", customerController.DeactivateCustomer)
}
