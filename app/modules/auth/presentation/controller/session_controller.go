package controller

import (
	cmd "clinic-vet-api/app/modules/auth/application/command"
	cmdBus "clinic-vet-api/app/modules/auth/infrastructure/bus"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TokenController struct {
	validator *validator.Validate
	bus       cmdBus.AuthCommandBus
}

func NewTokenController(validator *validator.Validate, bus cmdBus.AuthCommandBus) *TokenController {
	return &TokenController{
		validator: validator,
		bus:       bus,
	}
}

func (ctrl *TokenController) RefreshSession(c *gin.Context) {
	refreshToken, exists := c.Get("jwtToken")
	if !exists || refreshToken == "" {
		response.BadRequest(c, autherror.MissingRefreshTokenError())
		return
	}

	command := cmd.NewRefreshSessionCommand(refreshToken.(string))
	result := ctrl.bus.RefreshSession(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Session(), result.Message())
}

func (ctrl *TokenController) RevokeToken(c *gin.Context) {
}

func (ctrl *TokenController) RevokeAllTokens(c *gin.Context) {
}
