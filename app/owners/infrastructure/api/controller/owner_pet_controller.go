package ownerController

/*
import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OwnerPetController struct {
	ownerService services.OwnerService
	petService   services.PetService
	validator    *validator.Validate
}

func NewOwnerPetController(ownerService services.OwnerService, petService services.PetService) *OwnerPetController {
	return &OwnerPetController{
		ownerService: ownerService,
		petService:   petService,
		validator:    validator.New(),
	}
}

func (opc OwnerPetController) AddPet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var petInsertDTO DTOs.PetInsertDTO

		if err := c.BodyParser(&petInsertDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := opc.validator.Struct(petInsertDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		ownerDTO, err := opc.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		if err := opc.petService.CreatePet(petInsertDTO, ownerDTO.Id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(responses.SuccessResponse{
			Message: "Pet Successfully Created.",
		})
	}
}

func (opc OwnerPetController) GetMyPets() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		ownerDTO, err := opc.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		petListDTO, err := opc.petService.GetPetsByOwnerID(ownerDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Can't Get Pets",
			})
		}

		return c.Status(fiber.StatusOK).JSON(petListDTO)
	}
}

func (opc OwnerPetController) UpdatePet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var petUpdateDTO DTOs.PetUpdateDTO

		if err := c.BodyParser(&petUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		if err := opc.validator.Struct(&petUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		_, err := opc.petService.GetPetById(petUpdateDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		ownerDTO, err := opc.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		if err := opc.petService.UpdatePet(petUpdateDTO, ownerDTO.Id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Pet Successfully Updated.",
		})
	}
}

func (opc OwnerPetController) DeletePet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		petId, err := utils.ParseParamToInt(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		ownerDTO, err := opc.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		petDTO, err := opc.petService.GetPetById(petId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})

		}

		if ownerDTO.Id != petDTO.OwnerID {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: "Unauthorized Pet Delete.",
			})
		}

		if err := opc.petService.DeletePetById(petId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Pet Successfully Deleted.",
		})
	}
}
*/
