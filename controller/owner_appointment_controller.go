package controller

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OwnerAppointmentController struct {
	appointmentService services.ClientAppointmentService
	ownerService       services.OwnerService
	validator          *validator.Validate
}

func NewOwnerAppointmentController(appointmentService services.ClientAppointmentService, ownerService services.OwnerService) *OwnerAppointmentController {
	return &OwnerAppointmentController{
		appointmentService: appointmentService,
		ownerService:       ownerService,
		validator:          validator.New(),
	}
}

func (cac OwnerAppointmentController) RequestAnAppointment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var appointmentInsertDTO DTOs.AppointmentInsertDTO

		if err := c.BodyParser(&appointmentInsertDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := cac.validator.Struct(appointmentInsertDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		ownerDTO, err := cac.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		appointmentDTO, err := cac.appointmentService.RequestAnAppointment(appointmentInsertDTO, ownerDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(appointmentDTO)
	}
}

func (cac OwnerAppointmentController) GetMyAppointments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		ownerDTO, err := cac.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		appointments, err := cac.appointmentService.GetAppointmentByOwnerId(ownerDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(appointments)
	}
}

func (cac OwnerAppointmentController) UpdateAnAppointment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var appointmentUpdateDTO DTOs.AppointmentUpdateDTO

		if err := c.BodyParser(&appointmentUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := cac.validator.Struct(appointmentUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		ownerDTO, err := cac.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		err = cac.appointmentService.UpdateAppointment(appointmentUpdateDTO, ownerDTO.Id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Can't Update Appointment",
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Appointment Successfully Updated.",
		})
	}
}

func (cac OwnerAppointmentController) CancelAnAppointment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		appointmentID, err := utils.ParseID(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		userID, err := utils.GetUserIDFromRequest(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		ownerDTO, err := cac.ownerService.GetOwnerByUserID(int32(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Owner Not Found",
			})
		}

		appointmentDTO, err := cac.appointmentService.GetAppointmentById(appointmentID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Appointment Not Found",
			})
		}

		if appointmentDTO.OwnerID != ownerDTO.Id {
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
					Message: "Cancellation Not Allowed",
				})
			}
		}

		if err := cac.appointmentService.CancelAppointmentById(appointmentID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Appointment Successfully Deleted.",
		})
	}
}
