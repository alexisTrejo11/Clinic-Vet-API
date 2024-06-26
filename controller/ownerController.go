package controller

import (
	"strconv"

	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OwnerController struct {
	ownerService *services.OwnerService
	validator    *validator.Validate
}

func NewOwnerRepository(ownerService *services.OwnerService) *OwnerController {
	return &OwnerController{
		ownerService: ownerService,
		validator:    validator.New(),
	}
}

func (oc *OwnerController) CreateOwner() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newOwner dtos.OwnerInsertDTO

		if err := c.BodyParser(&newOwner); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid JSON payload",
			})
		}

		if err := oc.validator.Struct(newOwner); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Validation failed",
				"error":   err.Error(),
			})
		}

		if err := oc.ownerService.CreateOwner(&newOwner); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Can Not Create Owner",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"created": "Owner Created Successfully Created!.",
		})
	}

}

func (oc *OwnerController) GetOwnerById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Read Param
		ownerIdStr := c.Params("id")
		if ownerIdStr == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Id is Empty",
			})
		}

		// Parse Param
		intValue, err := strconv.Atoi(ownerIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Cant Proccess Id",
			})
		}

		onwerId := int32(intValue)

		ownerDTO, err := oc.ownerService.GetOwnerById(onwerId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Owner Not Found",
			})

		}
		return c.Status(fiber.StatusOK).JSON(ownerDTO)
	}
}
