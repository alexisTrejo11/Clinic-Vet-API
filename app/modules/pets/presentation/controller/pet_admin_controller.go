package controller

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/pets/application/usecase"
	"clinic-vet-api/app/modules/pets/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetController struct {
	validator   *validator.Validate
	petUseCases usecase.PetUseCases
}

func NewPetController(validator *validator.Validate, petUseCases usecase.PetUseCases) *PetController {
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
func (c *PetController) SearchPets(ctx *gin.Context) {
	var searchParams dto.PetSearchRequest
	if err := ginUtils.BindAndValidateQuery(ctx, &searchParams, c.validator); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	searchSpecification := searchParams.ToSpecification()
	pets, err := c.petUseCases.SearchPets(ctx, *searchSpecification)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Found(ctx, pets, "Pets")
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
	id, err := ginUtils.ParseParamToUInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "pet_id", ctx.Param("id")))
		return
	}

	pet, err := c.petUseCases.FindPetByID(ctx, valueobject.NewPetID(id))
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Found(ctx, pet, "Pet")
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
	var requestData dto.AdminCreatePetRequest
	if err := ginUtils.BindAndValidateBody(ctx, &requestData, c.validator); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	createPetData := requestData.ToCreateData()
	pet, err := c.petUseCases.CreatePet(ctx.Request.Context(), createPetData)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Created(ctx, pet.ID, "Pet")
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
	id, err := ginUtils.ParseParamToUInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "pet_id", ctx.Param("id")))
		return
	}

	var requestData dto.AdminUpdatePetRequest
	if err := ginUtils.BindAndValidateBody(ctx, &requestData, c.validator); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	updatePetData := requestData.ToUpdatePet(id)
	_, err = c.petUseCases.UpdatePet(ctx, *updatePetData)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Updated(ctx, nil, "Pet")
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
	id, err := ginUtils.ParseParamToUInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "pet_id", ctx.Param("id")))
		return
	}

	petID := valueobject.NewPetID(id)
	if err := c.petUseCases.DeletePet(ctx, petID, true); err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.NoContent(ctx)
}
