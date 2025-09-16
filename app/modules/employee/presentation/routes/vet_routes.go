package routes

import (
	"clinic-vet-api/app/modules/employee/presentation/controller"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine, employeeController *controller.EmployeeController) {
	v2Router := router.Group("/api/v2/employees")
	v2Router.GET("/:id", employeeController.GetEmployeeByID)
	v2Router.GET("/", employeeController.SearchEmployees)
	v2Router.POST("/", employeeController.CreateEmployee)
	v2Router.PATCH("/:id", employeeController.UpdateEmployee)
	v2Router.DELETE("/:id", employeeController.DeleteEmployee)
}
