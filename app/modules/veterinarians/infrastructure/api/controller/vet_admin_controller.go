package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/usecase"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/infrastructure/api/dto"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VeterinarianController struct {
	validator   *validator.Validate
	vetUseCases usecase.VeterinarianUseCases
}

func NewVeterinarianController(
	validator *validator.Validate,
	vetUseCases usecase.VeterinarianUseCases,
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
// @Success 200 {array} dto.VetResponse "List of veterinarians"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /veterinarians [get]
func (controller *VeterinarianController) ListVeterinarians(c *gin.Context) {
	listParams, err := fetchVetParams(c)
	if err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	vets, err := controller.vetUseCases.ListVetUseCase(c.Request.Context(), listParams)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, vets)
}

// @Summary Get a veterinarian by ID
// @Description Retrieves a single veterinarian by their unique ID.
// @Tags Veterinarians
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Success 200 {object} dto.VetResponse "Veterinarian details"
// @Failure 400 {object} response.APIResponse "Invalid ID supplied"
// @Failure 404 {object} response.APIResponse "Veterinarian not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /veterinarians/{id} [get]
func (controller *VeterinarianController) GetVeterinarianByID(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	veterinarian, err := controller.vetUseCases.GetVetByIDUseCase(c.Request.Context(), id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, veterinarian)
}

// @Summary Create a new veterinarian
// @Description Creates a new veterinarian with the provided details.
// @Tags Veterinarians
// @Accept json
// @Produce json
// @Param vetCreate body dto.VetCreate true "Veterinarian data"
// @Success 201 {object} dto.VetResponse "Veterinarian created"
// @Failure 400 {object} response.APIResponse "Invalid request body"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /veterinarians [post]
func (controller *VeterinarianController) CreateVeterinarian(c *gin.Context) {
	var requestData dto.CreateVetRequest

	if err := c.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createVetData, err := requestData.ToCreateData()
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	vetCreated, err := controller.vetUseCases.CreateVetUseCase(c.Request.Context(), createVetData)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, vetCreated)
}

// @Summary Update an existing veterinarian
// @Description Updates the details of an existing veterinarian by ID.
// @Tags Veterinarians
// @Accept json
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Param requestData body dto.VetUpdate true "Updated veterinarian data"
// @Success 200 {object} dto.VetResponse "Veterinarian updated"
// @Failure 400 {object} response.APIResponse "Invalid request data or ID"
// @Failure 404 {object} response.APIResponse "Veterinarian not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /veterinarians/{id} [patch]
func (controller *VeterinarianController) UpdateVeterinarian(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "Veterinarian_id", c.Param("id")))
		return
	}

	var requestData dto.UpdateVetRequest
	if err := c.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	updateVetData, err := requestData.ToUpdateData(id)
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	verUpdated, err := controller.vetUseCases.UpdateVetUseCase(context.TODO(), updateVetData)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, verUpdated)
}

// @Summary Delete a veterinarian
// @Description Deletes a veterinarian from the system by ID.
// @Tags Veterinarians
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Success 204 "No Content"
// @Failure 400 {object} response.APIResponse "Invalid ID supplied"
// @Failure 404 {object} response.APIResponse "Veterinarian not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /veterinarians/{id} [delete]
func (controller *VeterinarianController) DeleteVeterinarian(c *gin.Context) {
	id, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "Veterinarian_id", c.Param("id")))
		return
	}

	if err := controller.vetUseCases.DeleteVet(context.TODO(), id); err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.NoContent(c)
}
