package controller

import (
	"clinic-vet-api/app/modules/pets/presentation/dto"
	"clinic-vet-api/app/modules/pets/presentation/service"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetController struct {
	validator  *validator.Validate
	operations *service.PetControllerOperations
}

func NewPetController(validator *validator.Validate, operations *service.PetControllerOperations) *PetController {
	return &PetController{
		validator:  validator,
		operations: operations,
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
func (ctrl *PetController) SearchPets(c *gin.Context) {
	ctrl.operations.FindPetsBySpecification(c)
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
func (ctrl *PetController) FindPetByID(c *gin.Context) {
	ctrl.operations.FindPetByID(c, nil)
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
func (ctrl *PetController) CreatePet(c *gin.Context) {
	var requestData dto.AdminPetCreateExtraFields
	if err := ginUtils.ShouldBindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}
	ctrl.operations.CreatePet(c, requestData.CustomerID, requestData.IsActive)
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
func (ctrl *PetController) UpdatePet(c *gin.Context) {
	var extraFields dto.AdminUpdatePetExtraFields
	if err := ginUtils.ShouldBindAndValidateBody(c, &extraFields, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	ctrl.operations.UpdatePet(c, extraFields.CustomerID, extraFields.IsActive)
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
func (ctrl *PetController) DeletePet(c *gin.Context) {
	hardDelete := c.Query("hard")
	hardDelete = strings.ToLower(hardDelete)
	hardDelete = strings.TrimSpace(hardDelete)

	if hardDelete != "true" && hardDelete != "false" {
		err := httpError.RequestURLQueryError(fmt.Errorf("invalid value for 'hard' query parameter: %s. must be 'true' or 'false'", hardDelete), c.Request.URL.String())
		response.BadRequest(c, err)
		return
	}

	ctrl.operations.DeletePet(c, nil, hardDelete == "true")
}
