package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/usecase"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerController struct {
	validator   *validator.Validate
	ownerFacade usecase.OwnerServiceFacade
}

func NewOwnerController(
	validator *validator.Validate,
	ownerFacade usecase.OwnerServiceFacade,
) *OwnerController {
	return &OwnerController{
		validator:   validator,
		ownerFacade: ownerFacade,
	}
}

func (c *OwnerController) ListOwners(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	status := ctx.Query("status")
	include_pets := ctx.Query("include_pets")

	searchParams, err := dto.NewOwnerSearch(limit, offset, status, include_pets)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	owners, err := c.ownerFacade.ListOwners(context.TODO(), *searchParams)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, owners)
}

func (c *OwnerController) GetOwnerByID(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "owner_id", ctx.Param("id"))
		return
	}

	ownerID, _ := valueobject.NewOwnerID(id)
	owner, err := c.ownerFacade.GetOwnerByID(context.TODO(), ownerID)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, owner)
}

func (c *OwnerController) CreateOwner(ctx *gin.Context) {
	var ownerCreate dto.OwnerCreate

	if err := ctx.ShouldBindBodyWithJSON(&ownerCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&ownerCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	owner, err := c.ownerFacade.CreateOwner(context.TODO(), ownerCreate)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Created(ctx, owner)
}

func (c *OwnerController) UpdateOwner(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "owner_id", ctx.Param("id"))
		return
	}

	var ownerUpdate dto.OwnerUpdate
	if err := ctx.ShouldBindBodyWithJSON(&ownerUpdate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&ownerUpdate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	ownerID, _ := valueobject.NewOwnerID(id)
	Owner, err := c.ownerFacade.UpdateOwner(context.TODO(), ownerID, ownerUpdate)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, Owner)
}

func (c *OwnerController) DeleteOwner(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "owner_id", ctx.Param("id"))
		return
	}

	ownerID, _ := valueobject.NewOwnerID(id)
	if err := c.ownerFacade.SoftDeleteOwner(context.TODO(), ownerID); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
