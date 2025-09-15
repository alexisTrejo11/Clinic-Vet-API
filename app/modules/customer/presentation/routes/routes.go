// Package routes defines the HTTP routes for customer-related operations.
package routes

import (
	"clinic-vet-api/app/modules/customer/presentation/controller"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(app *gin.Engine, customerController *controller.CustomerController) {
	customerV2 := app.Group("/api/v2/customer")
	customerV2.GET("/:id", customerController.GetCustomerByID)
	customerV2.GET("/", customerController.SearchCustomers)
	customerV2.POST("/", customerController.CreateCustomer)
	customerV2.PATCH("/:id", customerController.UpdateCustomer)
	customerV2.DELETE("/:id", customerController.DeactivateCustomer)
}
