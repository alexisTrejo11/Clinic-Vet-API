package controller

import (
	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OwnerController struct {
	ownerService services.OwnerService
	validator    *validator.Validate
	logger       *logrus.Logger
}

func NewOwnerController(ownerService services.OwnerService) *OwnerController {
	return &OwnerController{
		ownerService: ownerService,
		validator:    validator.New(),
		logger:       logrus.New(),
	}
}

// CreateOwner godoc
// @Summary Create a new owner
// @Description Create a new owner with the input payload
// @Tags owners
// @Accept  json
// @Produce  json
// @Param owner body dtos.OwnerInsertDTO true "Owner to create"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /owners [post]
func (oc *OwnerController) CreateOwner() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var owner dtos.OwnerInsertDTO

		if err := c.BodyParser(&owner); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := oc.validator.Struct(owner); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		if err := oc.ownerService.CreateOwner(&owner); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Cannot create owner",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(responses.SuccessResponse{
			Message: "Owner created successfully",
		})
	}
}

// GetOwnerById godoc
// @Summary Get an owner by ID
// @Description Get an owner by their ID
// @Tags owners
// @Accept  json
// @Produce  json
// @Param id path int true "Owner ID"
// @Success 200 {object} dtos.OwnerReturnDTO
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /owners/{id} [get]
func (oc *OwnerController) GetOwnerById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ownerID, err := utils.ParseID(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid ID",
			})
		}

		ownerDTO, err := oc.ownerService.GetOwnerById(ownerID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner not found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(ownerDTO)
	}
}

// UpdateOwner godoc
// @Summary Update an owner
// @Description Update an owner with the input payload
// @Tags owners
// @Accept  json
// @Produce  json
// @Param owner body dtos.OwnerUpdateDTO true "Owner to update"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /owners [put]
func (oc *OwnerController) UpdateOwner() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ownerUpdateDTO dtos.OwnerUpdateDTO

		if err := c.BodyParser(&ownerUpdateDTO); err != nil {
			oc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse request body for update")

			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := oc.validator.Struct(ownerUpdateDTO); err != nil {
			oc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for owner update")

			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   "Id is obligatory",
			})
		}

		isOwnerExists := oc.ownerService.ValidateExistingOwner(ownerUpdateDTO.Id)
		if !isOwnerExists {
			oc.logger.WithFields(logrus.Fields{"id": ownerUpdateDTO.Id}).Warn("Owner does not exist")

			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Doesn't Exist",
			})
		}

		if err := oc.ownerService.UpdateOwner(&ownerUpdateDTO); err != nil {
			oc.logger.WithFields(logrus.Fields{"id": ownerUpdateDTO.Id, "error": err.Error()}).Error("Failed to update owner")

			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
			})
		}

		oc.logger.WithFields(logrus.Fields{
			"id": ownerUpdateDTO.Id,
		}).Info("Owner successfully updated")

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Owner successfully updated",
		})
	}
}

// DeleteOwner godoc
// @Summary Delete an owner
// @Description Delete an owner by ID
// @Tags owners
// @Accept  json
// @Produce  json
// @Param id path int true "Owner ID"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /owners/{id} [delete]
func (oc *OwnerController) DeleteOwner() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ownerID, err := utils.ParseID(c)
		if err != nil {
			oc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse owner ID for deletion")

			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid ID",
			})
		}

		isOwnerExists := oc.ownerService.ValidateExistingOwner(ownerID)
		if !isOwnerExists {
			oc.logger.WithFields(logrus.Fields{"id": ownerID}).Warn("Owner does not exist")

			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Doesn't Exist",
			})
		}

		if err := oc.ownerService.DeleteOwner(ownerID); err != nil {
			oc.logger.WithFields(logrus.Fields{"id": ownerID, "error": err.Error()}).Error("Failed to delete owner")

			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
			})
		}

		oc.logger.WithFields(logrus.Fields{"id": ownerID}).Info("Owner successfully deleted")

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Owner successfully deleted",
		})
	}
}
