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

// This package contains wrapper functions that hold the Swagger annotations.
// They then delegate the actual work to the controller to keep it clean.

// @Summary List all veterinarians
// @Description Retrieves a list of all veterinarians with optional filtering and pagination.
// @Tags Veterinarians
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items per page"
// @Param specialty query string false "Filter by specialty"
// @Success 200 {array} vetDtos.VetResponse "List of veterinarians"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /veterinarians [get]
func (c *VeterinarianController) ListVeterinarians(ctx *gin.Context) {
	listParams, err := fetchVetParams(ctx)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	vets, err := c.vetUseCases.ListVetUseCase(ctx.Request.Context(), listParams)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, vets)
}

// @Summary Get a veterinarian by ID
// @Description Retrieves a single veterinarian by their unique ID.
// @Tags Veterinarians
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Success 200 {object} vetDtos.VetResponse "Veterinarian details"
// @Failure 400 {object} apiResponse.APIResponse "Invalid ID supplied"
// @Failure 404 {object} apiResponse.APIResponse "Veterinarian not found"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /veterinarians/{id} [get]
func (c *VeterinarianController) GetVeterinarianById(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "Veterinarian_id", ctx.Param("id"))
		return
	}

	Veterinarian, err := c.vetUseCases.GetVetByIdUseCase(context.TODO(), id)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, Veterinarian)
}

// @Summary Create a new veterinarian
// @Description Creates a new veterinarian with the provided details.
// @Tags Veterinarians
// @Accept json
// @Produce json
// @Param vetCreate body vetDtos.VetCreate true "Veterinarian data"
// @Success 201 {object} vetDtos.VetResponse "Veterinarian created"
// @Failure 400 {object} apiResponse.APIResponse "Invalid request body"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /veterinarians [post]
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
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Created(ctx, vetCreated)
}

// @Summary Update an existing veterinarian
// @Description Updates the details of an existing veterinarian by ID.
// @Tags Veterinarians
// @Accept json
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Param vetUpdate body vetDtos.VetUpdate true "Updated veterinarian data"
// @Success 200 {object} vetDtos.VetResponse "Veterinarian updated"
// @Failure 400 {object} apiResponse.APIResponse "Invalid request data or ID"
// @Failure 404 {object} apiResponse.APIResponse "Veterinarian not found"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /veterinarians/{id} [patch]
func (c *VeterinarianController) UpdateVeterinarian(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
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
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.Success(ctx, verUpdated)
}

// @Summary Delete a veterinarian
// @Description Deletes a veterinarian from the system by ID.
// @Tags Veterinarians
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Success 204 "No Content"
// @Failure 400 {object} apiResponse.APIResponse "Invalid ID supplied"
// @Failure 404 {object} apiResponse.APIResponse "Veterinarian not found"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /veterinarians/{id} [delete]
func (c *VeterinarianController) DeleteVeterinarian(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "Veterinarian_id", ctx.Param("id"))
		return
	}

	if err := c.vetUseCases.DeleteVetUseCase(context.TODO(), id); err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
