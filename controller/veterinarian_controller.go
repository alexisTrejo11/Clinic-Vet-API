package controller

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type VeterinarianController struct {
	validator  *validator.Validate
	logger     *logrus.Logger
	vetService services.VeterinarianService
}

func NewVeterinarianController(vetService services.VeterinarianService) *VeterinarianController {
	return &VeterinarianController{
		vetService: vetService,
		validator:  validator.New(),
		logger:     logrus.New(),
	}
}

// CreateVeterinarian godoc
// @Summary Create a new veterinarian
// @Description Create a new veterinarian with the input payload
// @Tags veterinarians
// @Accept  json
// @Produce  json
// @Param vet body DTOs.VetInsertDTO true "Veterinarian to create"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /veterinarians [post]
func (vc VeterinarianController) CreateVeterinarian() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var veterinarianInsertDTO DTOs.VetInsertDTO

		if err := c.BodyParser(&veterinarianInsertDTO); err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse request body")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := vc.validator.Struct(veterinarianInsertDTO); err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for veterinarian creation")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		if err := vc.vetService.CreateVeterinarian(veterinarianInsertDTO); err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to create veterinarian")
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Cannot create veterinarian",
			})
		}

		vc.logger.Info("Veterinarian created successfully")
		return c.Status(fiber.StatusCreated).JSON(responses.SuccessResponse{
			Message: "Veterinarian created successfully",
		})
	}
}

// GetVeterinarianById godoc
// @Summary Get a veterinarian by ID
// @Description Get a veterinarian by their ID
// @Tags veterinarians
// @Accept  json
// @Produce  json
// @Param id path int true "Veterinarian ID"
// @Success 200 {object} DTOs.VetReturnDTO
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /veterinarians/{id} [get]
func (vc VeterinarianController) GetVeterinarianById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		veterinarianID, err := utils.ParseID(c)
		if err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Invalid ID for veterinarian")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid ID",
			})
		}

		vetDTO, err := vc.vetService.GetVeterinarianById(veterinarianID)
		if err != nil {
			vc.logger.WithFields(logrus.Fields{"id": veterinarianID}).Warn("Veterinarian not found")
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Veterinarian not found",
			})
		}

		vc.logger.WithFields(logrus.Fields{"id": veterinarianID}).Info("Veterinarian retrieved successfully")
		return c.Status(fiber.StatusOK).JSON(vetDTO)
	}
}

// UpdateVeterinarian godoc
// @Summary Update an existing veterinarian
// @Description Update a veterinarian's information with the input payload
// @Tags veterinarians
// @Accept  json
// @Produce  json
// @Param vet body DTOs.VetUpdateDTO true "Veterinarian to update"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /veterinarians [put]
func (vc VeterinarianController) UpdateVeterinarian() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var veterinarianUpdateDTO DTOs.VetUpdateDTO

		if err := c.BodyParser(&veterinarianUpdateDTO); err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse request body for update")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := vc.validator.Struct(veterinarianUpdateDTO); err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for veterinarian update")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   "Id is obligatory",
			})
		}

		isVetExisting := vc.vetService.ValidateExistingVet(veterinarianUpdateDTO.Id)
		if !isVetExisting {
			vc.logger.WithFields(logrus.Fields{"id": veterinarianUpdateDTO.Id}).Warn("Veterinarian does not exist")
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Veterinarian Doesn't Exist",
			})
		}

		if err := vc.vetService.UpdateVeterinarian(veterinarianUpdateDTO); err != nil {
			vc.logger.WithFields(logrus.Fields{"id": veterinarianUpdateDTO.Id, "error": err.Error()}).Error("Failed to update veterinarian")
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
			})
		}

		vc.logger.WithFields(logrus.Fields{"id": veterinarianUpdateDTO.Id}).Info("Veterinarian updated successfully")
		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Veterinarian updated successfully",
		})
	}
}

// DeleteVeterinarian godoc
// @Summary Delete a veterinarian by ID
// @Description Delete a veterinarian by their ID
// @Tags veterinarians
// @Accept  json
// @Produce  json
// @Param id path int true "Veterinarian ID"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /veterinarians/{id} [delete]
func (vc VeterinarianController) DeleteVeterinarian() fiber.Handler {
	return func(c *fiber.Ctx) error {
		vetID, err := utils.ParseID(c)
		if err != nil {
			vc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Invalid ID for veterinarian deletion")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid ID",
			})
		}

		isVetExisting := vc.vetService.ValidateExistingVet(vetID)
		if !isVetExisting {
			vc.logger.WithFields(logrus.Fields{"id": vetID}).Warn("Veterinarian does not exist")
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Veterinarian Doesn't Exist",
			})
		}

		if err := vc.vetService.DeleteVeterinarianbyId(vetID); err != nil {
			vc.logger.WithFields(logrus.Fields{"id": vetID, "error": err.Error()}).Error("Failed to delete veterinarian")
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
			})
		}

		vc.logger.WithFields(logrus.Fields{"id": vetID}).Info("Veterinarian successfully deleted")
		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Veterinarian successfully deleted",
		})
	}
}
