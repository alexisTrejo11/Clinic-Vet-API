package controller

/*
import (
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ClientMedicalHistory struct {
	petServices           services.PetService
	medicalHistoryService services.MedicalHistoryService
	customerService          services.customerService
	validator             *validator.Validate
}

func NewClientMedicalHistory(petServices services.PetService, customerServices services.customerService, medicalHistoryService services.MedicalHistoryService) *ClientMedicalHistory {
	return &ClientMedicalHistory{
		petServices:           petServices,
		medicalHistoryService: medicalHistoryService,
		customerService:          customerServices,
		validator:             validator.New(),
	}
}

func (cmh ClientMedicalHistory) GetMyPetsMedicalHistories() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
				Error:   "Can't Parse UserID",
			})
		}

		customerDTO, err := cmh.customerService.GetcustomerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "customer Not Found",
			})
		}

		medicalHistoriesDTOs, err := cmh.medicalHistoryService.GetMedicalRepositoryBycustomer(*customerDTO)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(medicalHistoriesDTOs)
	}
}

func (cmh ClientMedicalHistory) GetMyPetsMedicalHistoryByPetID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		petID, err := utils.ParseParamToInt(c)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
				Error:   "Can't Parse UserID",
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Server Error",
				Error:   "Can't Parse UserID",
			})
		}

		customerDTO, err := cmh.customerService.GetcustomerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "customer Not Found",
			})
		}

		medicalHistoriesDTOs, err := cmh.medicalHistoryService.GetMedicalRepositoryByPetID(petID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Pet Not Found",
			})
		}

		// Check If Medical Histories Belongs to Given customer ID
		isMedicalHistoriesValidated := cmh.petServices.ValidPetcustomer(petID, customerDTO.Id)
		if !isMedicalHistoriesValidated {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: "Unauthorized",
			})
		}

		return c.Status(fiber.StatusOK).JSON(medicalHistoriesDTOs)

	}
}
*/
