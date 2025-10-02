package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/medical/vaccination/presentation/controller"

	"github.com/gin-gonic/gin"
)

type VaccinationRoutes struct {
	clientController   *controller.CustomerPetVaccinationController
	employeeController *controller.EmployeePetVaccinationController
}

func NewVaccinationRoutes(clientController *controller.CustomerPetVaccinationController, employeeController *controller.EmployeePetVaccinationController) *VaccinationRoutes {
	return &VaccinationRoutes{
		clientController:   clientController,
		employeeController: employeeController,
	}
}

func (r *VaccinationRoutes) RegisterEmployeeRoutes(group *gin.RouterGroup, middleware *middleware.AuthMiddleware, controller *controller.EmployeePetVaccinationController) {
	employeeGroup := group.Group("employees/vaccinations")
	employeeGroup.Use(middleware.Authenticate())
	employeeGroup.Use(middleware.RequireAnyRole("employee, recepcionist, manager, veterinarian"))
	{
		employeeGroup.GET("/:id", controller.GetMyVaccinationAppliedDetail)
		employeeGroup.GET("", controller.GetMyVaccinationHistory)
		employeeGroup.GET("/pets/:petId", controller.GetMyVaccinationsHistoryByPet)
		employeeGroup.POST("", controller.RegisterNewVaccination)
		employeeGroup.PUT("/:id", controller.UpdateMyVaccinationApplied)
	}
}

func (r *VaccinationRoutes) RegisterCustomerRoutes(group *gin.RouterGroup, middleware *middleware.AuthMiddleware, controller *controller.CustomerPetVaccinationController) {
	customerGroup := group.Group("customers/pets")
	customerGroup.Use(middleware.Authenticate())
	customerGroup.Use(middleware.RequireAnyRole("customer"))
	{
		customerGroup.GET("/:id/vaccinations", controller.GetByMyPetVaccinationHistory)
		customerGroup.GET("/vaccinations", controller.GetByMyPetsVaccinationHistory)
		customerGroup.GET("/vaccinations/:id", controller.GetByMyPetVaccinationHistoryDetail)
	}
}
