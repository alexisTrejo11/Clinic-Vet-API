package controller

import (
	"strconv"

	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PetController struct {
	petService *services.PetService
	validator  *validator.Validate
}

const (
	badRequest    string = "Bad_Request"
	validationErr string = "Validation_Error"
	serverError   string = "Internal_Server_Error"
	notFound      string = "NOT_FOUND"
	created       string = "Created"
	ok            string = "OK"
	updateMsg     string = "Pet Succesfully Updated!."
	createMsg     string = "Pet Succesfully Created!."
	deleteMsg     string = "Pet Succesfully Deleted!."
)

func NewPetController(petService *services.PetService) *PetController {
	return &PetController{
		petService: petService,
		validator:  validator.New(),
	}
}

func (pc *PetController) CreatePet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newPet dtos.PetInsertDTO

		if err := c.BodyParser(&newPet); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				badRequest: err.Error(),
			})

		}

		if err := pc.validator.Struct(&newPet); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				validationErr: err.Error(),
			})
		}

		if err := pc.petService.CreatePet(newPet, 1); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				serverError: err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			created: createMsg,
		})
	}
}

func (pc *PetController) GetPetById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		petIdStr := c.Params("petId")
		if petIdStr == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "petId is Empty",
			})
		}

		intValue, err := strconv.Atoi(petIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Cant Proccess petId",
			})
		}

		petId := int32(intValue)

		petDTO, err := pc.petService.GetPetById(petId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				notFound: err.Error(),
			})

		}
		return c.Status(fiber.StatusOK).JSON(petDTO)
	}
}

func (pc *PetController) UpdatePet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var petUpdateDTO dtos.PetUpdateDTO

		if err := c.BodyParser(&petUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				badRequest: err.Error(),
			})
		}

		if err := pc.validator.Struct(&petUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				validationErr: err.Error(),
			})
		}

		_, err := pc.petService.GetPetById(petUpdateDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				notFound: err.Error(),
			})

		}

		if err := pc.petService.UpdatePet(petUpdateDTO, 1); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				serverError: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			ok: updateMsg,
		})
	}
}

func (pc *PetController) DeletePet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		petIdStr := c.Params("petId")
		if petIdStr == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "petId is Empty",
			})
		}

		intValue, err := strconv.Atoi(petIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Cant Proccess petId",
			})
		}

		petId := int32(intValue)

		_, err = pc.petService.GetPetById(petId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				notFound: err.Error(),
			})
		}

		if err := pc.petService.DeletePetById(petId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				serverError: err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			ok: deleteMsg,
		})
	}
}
