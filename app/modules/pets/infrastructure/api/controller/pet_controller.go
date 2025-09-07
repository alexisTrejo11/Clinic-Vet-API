package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/usecase"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetController struct {
	validator   *validator.Validate
	petUseCases usecase.PetUseCasesFacade
}

func NewPetController(
	validator *validator.Validate,
	petUseCases usecase.PetUseCasesFacade,
) *PetController {
	return &PetController{
		validator:   validator,
		petUseCases: petUseCases,
	}
}

// ListPets lists all pets.
// @Summary List all pets
// @Description Retrieves a list of all pet records.
// @Tags Pet Management
// @Produce json
// @Success 200 {array} PetResponse "List of pets"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /pets [get]
func (c *PetController) ListPets(ctx *gin.Context) {
	pets, err := c.petUseCases.ListPets(ctx)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, pets)
}

// GetPetByID retrieves a pet by its ID.
// @Summary Get a pet by ID
// @Description Retrieves a single pet record by its unique ID.
// @Tags Pet Management
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} PetResponse "Pet found"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Failure 404 {object} response.APIResponse "Pet not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /pets/{id} [get]
func (c *PetController) GetPetByID(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "pet_id", ctx.Param("id"))
		return
	}

	petID, err := valueobject.NewPetID(id)
	if err != nil {
		response.BadRequest(ctx, err)
		return
	}

	pet, err := c.petUseCases.GetPetByID(ctx, petID)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, pet)
}

// CreatePet creates a new pet record.
// @Summary Create a new pet
// @Description Creates a new pet record with the provided data.
// @Tags Pet Management
// @Accept json
// @Produce json
// @Param pet body PetInsertRequest true "Pet creation request"
// @Success 201 {object} PetResponse "Pet created successfully"
// @Failure 400 {object} response.APIResponse "Invalid request body or validation error"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /pets [post]
func (c *PetController) CreatePet(ctx *gin.Context) {
	var petCreate PetInsertRequest

	if err := ctx.ShouldBindBodyWithJSON(&petCreate); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&petCreate); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	createPet := requestToCreatePet(petCreate)
	pet, err := c.petUseCases.CreatePet(ctx, createPet)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Created(ctx, pet)
}

// UpdatePet updates an existing pet record.
// @Summary Update an existing pet
// @Description Updates a pet record by its ID with the provided data.
// @Tags Pet Management
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param pet body PetInsertRequest true "Pet update request"
// @Success 200 {object} PetResponse "Pet updated successfully"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter or request body"
// @Failure 404 {object} response.APIResponse "Pet not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /pets/{id} [put]
func (c *PetController) UpdatePet(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "pet_id", ctx.Param("id"))
		return
	}

	var petCreate PetInsertRequest
	if err := ctx.ShouldBindBodyWithJSON(&petCreate); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	petUpdate := requestToUpdatePet(petCreate, id)
	if err := c.validator.Struct(&petUpdate); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	pet, err := c.petUseCases.UpdatePet(ctx, petUpdate)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, pet)
}

// SoftDeletePet soft deletes a pet record.
// @Summary Soft delete a pet
// @Description Marks a pet record as inactive without permanently deleting it.
// @Tags Pet Management
// @Param id path int true "Pet ID"
// @Success 204 "No Content"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Failure 404 {object} response.APIResponse "Pet not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /pets/{id} [delete]
func (c *PetController) SoftDeletePet(ctx *gin.Context) {
	id, err := utils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "pet_id", ctx.Param("id"))
		return
	}

	petID, err := valueobject.NewPetID(id)
	if err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := c.petUseCases.DeletePet(ctx, petID, true); err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.NoContent(ctx)
}
