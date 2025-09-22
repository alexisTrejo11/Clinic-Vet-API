// Package service contains the implementation of services related to pet commands.
package service

import (
	"clinic-vet-api/app/modules/pets/application/cqrs"
	"clinic-vet-api/app/modules/pets/application/cqrs/command"
	"clinic-vet-api/app/modules/pets/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetControllerOperations struct {
	bus       cqrs.PetServiceBus
	validator *validator.Validate
}

func NewPetControllerOperations(bus cqrs.PetServiceBus, validator *validator.Validate) *PetControllerOperations {
	return &PetControllerOperations{bus: bus, validator: validator}
}

func (s *PetControllerOperations) CreatePet(c *gin.Context, customerID *uint, isActive bool) {
	var requestBodyData dto.PetRequestData
	if err := ginUtils.BindAndValidateBody(c, &requestBodyData, s.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command := requestBodyData.ToCommand(*customerID, isActive)
	result := s.bus.CreatePet(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "Pet")
}

func (s *PetControllerOperations) UpdatePet(c *gin.Context, customerID *uint, isActive *bool) {
	petID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var requestBodyData dto.UpdatePetRequest
	if err := ginUtils.BindAndValidateBody(c, &requestBodyData, s.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command := requestBodyData.ToCommand(petID, customerID, isActive)
	result := s.bus.UpdatePet(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, nil, "Pet")
}

func (s *PetControllerOperations) RestorePet(c *gin.Context) {
	petID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command := command.NewRestorePetCommand(petID)
	result := s.bus.RestorePet(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, nil, "Pet")
}

func (s *PetControllerOperations) DeactivatePet(c *gin.Context, customerID *uint) {
	petID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command := command.NewDeactivatePetCommand(petID, customerID)
	result := s.bus.DeactivatePet(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, nil, "Pet")
}

func (s *PetControllerOperations) DeletePet(c *gin.Context, customerID *uint, hardDelete bool) {
	petID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	// If customerID is provided (provided customerID means that operation is request by customer),
	// it's a customer request, so perform a soft delete
	command := command.NewDeletePetCommand(petID, customerID, customerID != nil)
	result := s.bus.DeletePet(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.NoContent(c)
}
