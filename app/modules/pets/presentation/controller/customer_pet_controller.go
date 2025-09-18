package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/pets/application/usecase"
	"clinic-vet-api/app/modules/pets/presentation/dto"
	autherror "clinic-vet-api/app/shared/error/auth"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerPetController struct {
	validate *validator.Validate
	usecases usecase.PetUseCases
}

func NewCustomerPetController(validate *validator.Validate, usecases usecase.PetUseCases) *CustomerPetController {
	return &CustomerPetController{
		validate: validate,
		usecases: usecases,
	}
}

func (ctrl *CustomerPetController) RegisterNewPet(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	var requestData *dto.CustomerCreatePetRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validate.Struct(requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	petCreateData := requestData.ToCreateData(user.CustomerID)
	_, err := ctrl.usecases.CreatePet(c.Request.Context(), petCreateData)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, nil, "Pet registered successfully")
}

func (ctrl *CustomerPetController) UpdateMyPet(c *gin.Context) {
	petIDUint, err := ginUtils.ParseParamToUInt(c, "pet_id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "pet_id", c.Param("pet_id")))
		return
	}

	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	var requestData *dto.CustomerUpdatePetRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validate.Struct(requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	petUpdateData := requestData.ToUpdatePet(petIDUint, user.CustomerID)
	_, err = ctrl.usecases.UpdatePet(c.Request.Context(), *petUpdateData)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, nil, "Pet updated successfully")
}

func (ctrl *CustomerPetController) GetMyPets(c *gin.Context) {
	var pagination page.PageInput
	if err := ginUtils.BindAndValidateBody(c, &pagination, ctrl.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	pagination.SetDefaultsFieldsIfEmpty()

	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	pets, err := ctrl.usecases.FindsPetByCustomerID(c.Request.Context(), user.CustomerID, pagination)

	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, pets, "Pets")
}

func (c *CustomerPetController) GetMyPetDetailsByID() {
}
