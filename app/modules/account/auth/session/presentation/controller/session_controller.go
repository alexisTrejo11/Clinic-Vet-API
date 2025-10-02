package controller

import (
	"clinic-vet-api/app/modules/account/auth/session/application"
	"clinic-vet-api/app/modules/account/auth/session/application/command"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SessionController struct {
	validator *validator.Validate
	service   application.SessionFacadeService
}

func NewSessionController(validator *validator.Validate, service application.SessionFacadeService) *SessionController {
	return &SessionController{
		validator: validator,
		service:   service,
	}
}

func (ctrl *SessionController) RefreshSession(c *gin.Context) {
	refreshToken, exists := c.Get("jwtToken")
	if !exists || refreshToken == "" {
		response.BadRequest(c, autherror.MissingRefreshTokenError())
		return
	}

	command := command.NewRefreshSessionCommand(refreshToken.(string))
	result := ctrl.service.RefreshSession(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.Session(), result.Message())
}

func (ctrl *SessionController) RevokeToken(c *gin.Context) {
}

func (ctrl *SessionController) RevokeAllTokens(c *gin.Context) {
}
