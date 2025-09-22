package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/employee/presentation/controller"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(appGroup *gin.RouterGroup, employeeController *controller.EmployeeController, authMiddleware *middleware.AuthMiddleware) {
	employeeGroup := appGroup.Group("/employees")

	// Public routes
	employeeGroup.GET("/:id", employeeController.GetEmployeeByID)
	employeeGroup.GET("/", employeeController.SearchEmployees)

	//employeeGroup.Use(authMiddleware.Authenticate())
	//employeeGroup.Use(authMiddleware.RequireAnyRole("admin", "manager"))

	// Protected routes
	employeeGroup.POST("/", employeeController.CreateEmployee)
	employeeGroup.PATCH("/:id", employeeController.UpdateEmployee)
	employeeGroup.DELETE("/:id", employeeController.DeleteEmployee)
}
