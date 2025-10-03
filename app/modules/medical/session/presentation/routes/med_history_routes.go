// Package routes defines the API routes for the medical module.
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/medical/session/presentation/controller"

	"github.com/gin-gonic/gin"
)

type MedicalSessionRoutes struct {
	AdminController    *controller.AdminMedicalSessionController
	CustomerController *controller.CustomerMedicalSessionController
	EmployeeController *controller.EmployeeMedicalSessionController
}

func NewMedicalSessionRoutes(adminController *controller.AdminMedicalSessionController,
	customerController *controller.CustomerMedicalSessionController,
	employeeController *controller.EmployeeMedicalSessionController,
) *MedicalSessionRoutes {
	return &MedicalSessionRoutes{
		AdminController:    adminController,
		CustomerController: customerController,
		EmployeeController: employeeController,
	}
}

func (r *MedicalSessionRoutes) RegisterAdminRoutes(routerGroup *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	routes := routerGroup.Group("/admin/medical-sessions")
	//routes.Use(authMiddleware.Authenticate())
	//routes.Use(authMiddleware.RequireAnyRole("admin", "superadmin"))

	routes.GET("/", r.AdminController.SearchMedSessions)
	routes.GET("/:id", r.AdminController.GetMedicalSessionByID)
	routes.GET("/today", r.AdminController.GetTodayMedSessions)
	routes.POST("/", r.AdminController.CreateMedicalSession)
	routes.DELETE("/:id", r.AdminController.SoftDeleteMedicalSession)
}

func (r *MedicalSessionRoutes) RegisterCustomerRoutes(routerGroup *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	routes := routerGroup.Group("/customer/medical-sessions")
	routes.Use(authMiddleware.Authenticate())
	routes.Use(authMiddleware.RequireAnyRole("customer"))

	routes.GET("/", r.CustomerController.GetMyPetSessions)
	routes.GET("/:id", r.CustomerController.GetMyPetSessionsByPetID)
}

func (r *MedicalSessionRoutes) RegisterEmployeeRoutes(routerGroup *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	routes := routerGroup.Group("/employee/medical-sessions")
	routes.Use(authMiddleware.Authenticate())
	routes.Use(authMiddleware.RequireAnyRole(enum.UserRoleVeterinarian.String(), enum.UserRoleReceptionist.String()))

	routes.GET("/", r.EmployeeController.GetMyMedicalSessions)
	routes.GET("/:id", r.EmployeeController.GetMyMedicalSessionByID)
	routes.POST("/", r.EmployeeController.RegisterMedicalSession)
}
