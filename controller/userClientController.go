package controller

import (
	DTOs "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthClientController struct {
	authService services.AuthService
	validator   *validator.Validate
	logger      *logrus.Logger
}

func NewAuthClientController(authService services.AuthService) *AuthClientController {
	return &AuthClientController{
		authService: authService,
		validator:   validator.New(),
		logger:      logrus.New(),
	}

}

func (usc AuthClientController) ClientSignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userSignUpDTO DTOs.UserSignUpDTO

		if err := c.BodyParser(&userSignUpDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := usc.validator.Struct(userSignUpDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for signup")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		if err := usc.authService.ValidateUniqueFields(userSignUpDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid Given Credentials",
				Error:   err.Error(),
			})
		}

		JWT, err := usc.authService.CompleteSignUp(userSignUpDTO)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Internal Server Error",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(responses.SuccessResponse{
			Message: JWT,
		})
	}
}

func (usc AuthClientController) ClientLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userLoginDTO DTOs.UserLoginDTO

		if err := c.BodyParser(&userLoginDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := usc.validator.Struct(userLoginDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

		userDTO, err := usc.authService.FindUser(userLoginDTO)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "User Not Found With Given Credentials",
				Error:   err.Error(),
			})
		}

		JWT, err := usc.authService.CompleteLogin(*userDTO)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Internal Server Erro",
				Error:   err.Error(),
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(responses.SuccessResponse{
			Message: JWT,
		})

	}
}
