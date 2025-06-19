package controller

/*
import (
	DTOs "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/services"
	"example.com/at/backend/api-vet/utils/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthClientController struct {
	authClientService services.AuthClientService
	authCommonService services.AuthCommonService
	validator         *validator.Validate
	logger            *logrus.Logger
}

func NewAuthClientController(authClientService services.AuthClientService, authCommonService services.AuthCommonService) *AuthClientController {
	return &AuthClientController{
		authClientService: authClientService,
		authCommonService: authCommonService,
		validator:         validator.New(),
		logger:            logrus.New(),
	}
}

// ClientSignUp godoc
// @Summary Register a new client
// @Description Register a new client with the provided details
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body DTOs.UserSignUpDTO true "User registration details"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /auth/signup [post]
func (usc AuthClientController) ClientSignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userSignUpDTO DTOs.UserSignUpDTO

		if err := c.BodyParser(&userSignUpDTO); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse request body")
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

		if err := usc.authCommonService.ValidateUniqueFields(userSignUpDTO.Email, userSignUpDTO.PhoneNumber); err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Unique field validation failed for signup")
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Invalid Given Credentials",
				Error:   err.Error(),
			})
		}

		JWT, err := usc.authClientService.CompleteSignUp(userSignUpDTO)
		if err != nil {
			usc.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to complete signup process")
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Internal Server Error",
				Error:   "Can't Complete Signup",
			})
		}

		usc.logger.WithFields(logrus.Fields{"user": userSignUpDTO.Email}).Info("User signed up successfully")
		return c.Status(fiber.StatusCreated).JSON(responses.SuccessResponse{
			Message: JWT,
		})
	}
}

// ClientLogin godoc
// @Summary Log in an existing client
// @Description Log in an existing client with the provided credentials
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body DTOs.UserLoginDTO true "User login credentials"
// @Success 202 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /auth/login [post]
func (usc AuthClientController) ClientLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userLoginDTO DTOs.UserLoginDTO

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
				Message: "User Not Found With Given Credentials",
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
*/
