package vetController

import (
	"context"

	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VeterinarianController struct {
	validator   *validator.Validate
	vetUseCases vetUsecase.VeterinarianUseCases
}

func NewVeterinarianController(
	validator *validator.Validate,
	vetUseCases vetUsecase.VeterinarianUseCases,
) *VeterinarianController {
	return &VeterinarianController{
		validator:   validator,
		vetUseCases: vetUseCases,
	}
}

func (c *VeterinarianController) ListVeterinarians(ctx *gin.Context) {
	//limit := ctx.Query("limit")
	//offset := ctx.Query("offset")
	//status := ctx.Query("status")
	//include_pets := ctx.Query("include_pets")

	vets, err := c.vetUseCases.ListVetUseCase(context.TODO(), 1, 10)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Ok(ctx, vets)
}

func (c *VeterinarianController) GetVeterinarianById(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "Veterinarian_id", ctx.Param("id"))
		return
	}

	Veterinarian, err := c.vetUseCases.GetVetByIdUseCase(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Ok(ctx, Veterinarian)
}

func (c *VeterinarianController) CreateVeterinarian(ctx *gin.Context) {
	var vetCreate vetDtos.VetCreate

	if err := ctx.ShouldBindBodyWithJSON(&vetCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&vetCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	vetCreated, err := c.vetUseCases.CreateVetUseCase(context.TODO(), vetCreate)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Created(ctx, vetCreated)
}

func (c *VeterinarianController) UpdateVeterinarian(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "Veterinarian_id", ctx.Param("id"))
		return
	}

	var vetUpdate vetDtos.VetUpdate
	if err := ctx.ShouldBindBodyWithJSON(&vetUpdate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&vetUpdate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	verUpdated, err := c.vetUseCases.UpdateVetUseCase(context.TODO(), id, vetUpdate)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Ok(ctx, verUpdated)
}

func (c *VeterinarianController) DeleteVeterinarian(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "Veterinarian_id", ctx.Param("id"))
		return
	}

	if err := c.vetUseCases.DeleteVetUseCase(context.TODO(), id); err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
