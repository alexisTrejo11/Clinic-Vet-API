package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/medical/deworm/presentation/controller"

	"github.com/gin-gonic/gin"
)

type DewormRoutes struct {
	adminController    *controller.AdminDewormController
	customerController *controller.CustomerPetDewormController
	employeeController *controller.EmployeeDewormController
}

func NewDewormRoutes(
	adminController *controller.AdminDewormController,
	customerController *controller.CustomerPetDewormController,
	employeeController *controller.EmployeeDewormController,
) *DewormRoutes {
	return &DewormRoutes{
		adminController:    adminController,
		customerController: customerController,
		employeeController: employeeController,
	}
}

func (r *DewormRoutes) RegisterAdminRoutes(group *gin.RouterGroup, middleware *middleware.AuthMiddleware) {
	adminGroup := group.Group("pets/deworms")
	//adminGroup.Use(middleware.Authenticate())
	//adminGroup.Use(middleware.RequireAnyRole("admin"))
	{
		adminGroup.GET("/:id", r.adminController.GetDewormByID)
		adminGroup.GET("", r.adminController.SearhDeworms)
		adminGroup.POST("", r.adminController.CreateDeworm)
		adminGroup.PUT("/:id", r.adminController.UpdateDeworm)
		adminGroup.DELETE("/:id", r.adminController.DeleteDeworm)
	}
}

func (r *DewormRoutes) RegisterEmployees(group *gin.RouterGroup, middleware *middleware.AuthMiddleware) {
	employeeGroup := group.Group("employees/deworms")
	employeeGroup.Use(middleware.Authenticate())
	employeeGroup.Use(middleware.RequireAnyRole("employee, recepcionist, manager, veterinarian"))
	{
		employeeGroup.GET("/:id", r.employeeController.GetMyDewormAppliedByID)
		employeeGroup.GET("", r.employeeController.GetMyDewormsApplied)
		employeeGroup.POST("", r.employeeController.RegisterNewDewormApplication)
		employeeGroup.PUT("/:id", r.employeeController.UpdateMyDewormApplied)
	}
}

func (r *DewormRoutes) RegisterCustomerRoutes(group *gin.RouterGroup, middleware *middleware.AuthMiddleware) {
	customerGroup := group.Group("customers/pets")
	customerGroup.Use(middleware.Authenticate())
	customerGroup.Use(middleware.RequireAnyRole("customer"))
	{
		customerGroup.GET("/:id/deworms", r.customerController.GetMyPetDewormHistoryByPetID)
		customerGroup.GET("/deworms", r.customerController.GetMyPetDewormHistory)
		customerGroup.GET("/deworms/:id", r.customerController.GetMyPetDewormDetailByID)
	}
}
