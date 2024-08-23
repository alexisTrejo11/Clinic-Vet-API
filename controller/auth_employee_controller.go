package controller

import (
	DTOs "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthEmployeeController struct {
	authEmployeeService services.AuthEmployeeService
	authCommonService   services.AuthCommonService
	veterinarianService services.VeterinarianService
	validator           *validator.Validate
	logger              *logrus.Logger
}

func NewAuthEmployeeController(authEmployeeService services.AuthEmployeeService,
	authCommonService services.AuthCommonService,
	veterinarianService services.VeterinarianService) *AuthEmployeeController {
	return &AuthEmployeeController{
		authEmployeeService: authEmployeeService,
		authCommonService:   authCommonService,
		veterinarianService: veterinarianService,
		validator:           validator.New(),
		logger:              logrus.New(),
	}
}

func (usc AuthEmployeeController) EmployeeSignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userEmployeeSignUpDTO DTOs.UserEmployeeSignUpDTO

		if err := c.BodyParser(&userEmployeeSignUpDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse request body")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := usc.validator.Struct(userEmployeeSignUpDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for signup")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		// For User Employee Creation is necessary a exsiting veterinarian
		existingVet, err := usc.veterinarianService.GetVeterinarianById(userEmployeeSignUpDTO.VeterinarianId)
		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{
				Message: "Employee validation failed",
				Error:   "Employee is not registered",
			})
		}

		if err := usc.authCommonService.ValidateUniqueFields(userEmployeeSignUpDTO.Email, userEmployeeSignUpDTO.PhoneNumber); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Unique field validation failed for signup")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid Given Credentials",
				Error:   err.Error(),
			})
		}

		JWT, err := usc.authEmployeeService.CompleteSignUp(userEmployeeSignUpDTO, *existingVet)
		if err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to complete signup process")
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Internal Server Error",
				Error:   err.Error(),
			})
		}

		usc.logger.WithFields(logrus.Fields{"user": userEmployeeSignUpDTO.Email}).Info("User signed up successfully")
		return c.Status(fiber.StatusCreated).JSON(JWT)

	}
}

func (usc AuthEmployeeController) EmployeeLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userLoginDTO DTOs.UserEmployeeLoginDTO

		if err := c.BodyParser(&userLoginDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse request body")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := usc.validator.Struct(userLoginDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for login")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		userDTO, err := usc.authCommonService.FindUserByEmailOrPhone(userLoginDTO.Email, userLoginDTO.PhoneNumber)
		if err != nil {
			usc.logger.WithFields(logrus.Fields{"user": userLoginDTO.Email, "error": err.Error()}).Warn("User not found during login")
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Employee Not Found With Given Credentials",
				Error:   err.Error(),
			})
		}

		JWT, err := usc.authCommonService.CompleteLogin(*userDTO)
		if err != nil {
			usc.logger.WithFields(logrus.Fields{"user": userLoginDTO.Email, "error": err.Error()}).Error("Failed to complete login process")
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Internal Server Error",
				Error:   err.Error(),
			})
		}

		usc.logger.WithFields(logrus.Fields{"user": userLoginDTO.Email}).Info("User logged in successfully")
		return c.Status(fiber.StatusAccepted).JSON(responses.SuccessResponse{
			Message: JWT,
		})
	}
}
