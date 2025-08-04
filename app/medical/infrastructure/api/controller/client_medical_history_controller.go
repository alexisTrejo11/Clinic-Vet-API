package med_hist_controller

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
	ownerService          services.OwnerService
	validator             *validator.Validate
}

func NewClientMedicalHistory(petServices services.PetService, ownerServices services.OwnerService, medicalHistoryService services.MedicalHistoryService) *ClientMedicalHistory {
	return &ClientMedicalHistory{
		petServices:           petServices,
		medicalHistoryService: medicalHistoryService,
		ownerService:          ownerServices,
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

		ownerDTO, err := cmh.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		medicalHistoriesDTOs, err := cmh.medicalHistoryService.GetMedicalRepositoryByOwner(*ownerDTO)
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
		petID, err := utils.ParseID(c)

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

		ownerDTO, err := cmh.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		medicalHistoriesDTOs, err := cmh.medicalHistoryService.GetMedicalRepositoryByPetID(petID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Pet Not Found",
			})
		}

		// Check If Medical Histories Belongs to Given Owner ID
		isMedicalHistoriesValidated := cmh.petServices.ValidPetOwner(petID, ownerDTO.Id)
		if !isMedicalHistoriesValidated {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: "Unauthorized",
			})
		}

		return c.Status(fiber.StatusOK).JSON(medicalHistoriesDTOs)

	}
}
*/
