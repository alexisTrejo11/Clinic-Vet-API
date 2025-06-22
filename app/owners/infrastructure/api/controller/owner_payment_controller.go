package controller

/*
import (
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OwnerPaymentController struct {
	paymentService services.PaymentService
	ownerService   services.OwnerService
	validator      *validator.Validate
}

func NewOwnerPaymentController(paymentService services.PaymentService, ownerService services.OwnerService) *OwnerPaymentController {
	return &OwnerPaymentController{
		paymentService: paymentService,
		ownerService:   ownerService,
		validator:      validator.New(),
	}
}

func (opc OwnerPaymentController) GetMySuccesfullyPayments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		ownerDTO, err := opc.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		paymentDTOs, err := opc.paymentService.GetSuccessPaymentsToBePaidByOwnerID(ownerDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Payment Not Found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(paymentDTOs)
	}
}

func (opc OwnerPaymentController) GetMyPaymentsToBePaid() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		ownerDTO, err := opc.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		paymentDTOs, err := opc.paymentService.GetSuccessPaymentsToBePaidByOwnerID(ownerDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Payments Can't  Be Fetched.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(paymentDTOs)
	}

}


func (OwnerPaymentController) PayAnAppointment() fiber.Handler {

}
*/
