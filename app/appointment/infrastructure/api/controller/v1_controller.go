package controller

import (
	DTOs "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AppointmentController struct {
	appointmentService services.AppointmentService
	ownerService       services.OwnerService
	validator          *validator.Validate
}

func NewAppointmentController(appointmentService services.AppointmentService) *AppointmentController {
	return &AppointmentController{
		appointmentService: appointmentService,
		validator:          validator.New(),
	}
}

func (ac AppointmentController) CreateAppointment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var appointmentInsertDTO DTOs.AppointmentInsertDTO

		if err := c.BodyParser(&appointmentInsertDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := ac.validator.Struct(appointmentInsertDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		err := ac.appointmentService.CreateAppointment(appointmentInsertDTO)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Can't Create Appointment",
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Appointment Successfully Created.",
		})
	}
}

func (ac AppointmentController) GetAppointmentById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		appointmentID, err := utils.ParseID(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Cant Parse ID",
			})
		}

		appointmentDTO, err := ac.appointmentService.GetAppointmentById(appointmentID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Cant Found Appointment",
			})
		}

		return c.Status(fiber.StatusOK).JSON(appointmentDTO)
	}
}

func (ac AppointmentController) UpdateAnAppointment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var appointmentUpdateDTO DTOs.AppointmentUpdateDTO

		if err := c.BodyParser(&appointmentUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := ac.validator.Struct(appointmentUpdateDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		err := ac.appointmentService.UpdateAppointment(appointmentUpdateDTO)
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

func (ac AppointmentController) DeleteAppointment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		appointmentID, err := utils.ParseID(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Cant Parse ID",
			})
		}

		if _, err := ac.appointmentService.GetAppointmentById(appointmentID); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Cant Found Appointment",
			})
		}

		if err := ac.appointmentService.DeleteAppointmentById(appointmentID); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Cant Delete Appointment",
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
			Message: "Appointment Deleted",
		})
	}
}
