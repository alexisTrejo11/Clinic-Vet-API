package controller

import (
	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserClientController struct {
	userService services.UserService
	validator   *validator.Validate
	logger      *logrus.Logger
}

func NewUserClientController(userService services.UserService) *UserClientController {
	return &UserClientController{
		userService: userService,
		validator:   validator.New(),
		logger:      logrus.New(),
	}

}

func (usc UserClientController) ClientSignup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userSignupDTO dtos.UserSignupDTO

		if err := c.BodyParser(&userSignupDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid JSON payload",
			})
		}

		if err := usc.validator.Struct(userSignupDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Validation failed for signup")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Validation failed",
				Error:   err.Error(),
			})
		}

	}
}

func (usc UserClientController) ClientLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}
