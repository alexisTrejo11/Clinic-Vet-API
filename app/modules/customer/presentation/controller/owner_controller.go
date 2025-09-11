// Package controller defines the HTTP controllers for the owners module.
package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/usecase"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerController struct {
	validator *validator.Validate
	usecases  usecase.OwnerServiceFacade
}

func NewOwnerController(
	validator *validator.Validate,
	usecases usecase.OwnerServiceFacade,
) *OwnerController {
	return &OwnerController{
		validator: validator,
		usecases:  usecases,
	}
}

func (controller *OwnerController) SearchOwners(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	status := c.Query("status")
	include_pets := c.Query("include_pets")

	searchParams, err := dto.NewOwnerSearch(limit, offset, status, include_pets)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	owners, err := controller.usecases.SearchOwners(context.TODO(), *searchParams)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, owners)
}

func (controller *OwnerController) GetOwnerByID(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "owner_id", c.Param("id")))
		return
	}

	ownerID, _ := valueobject.NewOwnerID(id)
	owner, err := controller.usecases.GetOwnerByID(context.TODO(), ownerID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, owner)
}

func (controller *OwnerController) CreateOwner(c *gin.Context) {
	var ownerCreate dto.OwnerCreate

	if err := c.ShouldBindBodyWithJSON(&ownerCreate); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&ownerCreate); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	owner, err := controller.usecases.CreateOwner(context.TODO(), ownerCreate)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, owner)
}

func (controller *OwnerController) UpdateOwner(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "owner_id", c.Param("id")))
		return
	}

	var ownerUpdate dto.OwnerUpdate
	if err := c.ShouldBindBodyWithJSON(&ownerUpdate); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&ownerUpdate); err != nil {
		response.ApplicationError(c, err)
		return
	}

	ownerID, _ := valueobject.NewOwnerID(id)
	Owner, err := controller.usecases.UpdateOwner(context.TODO(), ownerID, ownerUpdate)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, Owner)
}

func (controller *OwnerController) DeleteOwner(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "owner_id", c.Param("id")))
		return
	}

	ownerID, _ := valueobject.NewOwnerID(id)
	if err := controller.usecases.SoftDeleteOwner(context.TODO(), ownerID); err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.NoContent(c)
}
