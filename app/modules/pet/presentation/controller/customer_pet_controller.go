package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/pet/presentation/service"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerPetController struct {
	validate   *validator.Validate
	operations *service.PetControllerOperations
}

func NewCustomerPetController(validate *validator.Validate, operations *service.PetControllerOperations) *CustomerPetController {
	return &CustomerPetController{
		validate:   validate,
		operations: operations,
	}
}

func (ctrl *CustomerPetController) RegisterNewPet(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	customerID := user.CustomerID
	ctrl.operations.CreatePet(c, &customerID, true)
}

func (ctrl *CustomerPetController) UpdateMyPet(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.UpdatePet(c, &user.CustomerID, nil)
}

func (ctrl *CustomerPetController) GetMyPets(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.FindPetsByCustomerID(c, user.CustomerID)
}

func (ctrl *CustomerPetController) GetMyPetDetails(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.FindPetByID(c, &user.CustomerID)
}

func (ctrl *CustomerPetController) DeleteMyPet(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.DeletePet(c, &user.CustomerID, false)
}
