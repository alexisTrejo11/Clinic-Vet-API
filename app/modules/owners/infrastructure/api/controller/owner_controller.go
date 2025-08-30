package controller

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/usecase"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerController struct {
	validator   *validator.Validate
	ownerFacade ownerUsecase.OwnerServiceFacade
}

func NewOwnerController(
	validator *validator.Validate,
	ownerFacade ownerUsecase.OwnerServiceFacade,
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

	searchParams, err := ownerDTOs.NewOwnerSearch(limit, offset, status, include_pets)
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

func (c *OwnerController) GetOwnerById(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "owner_id", ctx.Param("id"))
		return
	}

	owner, err := c.ownerFacade.GetOwnerById(context.TODO(), id)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, owner)
}

func (c *OwnerController) CreateOwner(ctx *gin.Context) {
	var ownerCreate ownerDTOs.OwnerCreate

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

	var ownerUpdate ownerDTOs.OwnerUpdate
	if err := ctx.ShouldBindBodyWithJSON(&ownerUpdate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&ownerUpdate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	Owner, err := c.ownerFacade.UpdateOwner(context.TODO(), id, ownerUpdate)
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

	if err := c.ownerFacade.SoftDeleteOwner(context.TODO(), id); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
