package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/infrastructure/bus"
	autherror "clinic-vet-api/app/shared/error/auth"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"
	"errors"

	"github.com/gin-gonic/gin"
)

type TwoFAAuthController struct {
	twoFaBus bus.TwoFAAuthCommandBus
	authBus  bus.AuthCommandBus
}

func NewTwoFAAuthController(twoFaBus bus.TwoFAAuthCommandBus, authBus bus.AuthCommandBus) *TwoFAAuthController {
	return &TwoFAAuthController{
		twoFaBus: twoFaBus,
		authBus:  authBus,
	}
}

func (ctrl *TwoFAAuthController) Send2FAToken(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	cmd := command.NewSend2FATokenCommand(userID.Value())
	result := ctrl.twoFaBus.Send2FAToken(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *TwoFAAuthController) TwoFaLogin(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	code := c.Query("code")
	if code == "" {
		response.BadRequest(c, errors.New("code query parameter is required"))
		return
	}

	meta := ginutils.NewLoginMetadata(c)
	cmd, err := command.NewTwoFactorLoginCommand(
		userID,
		code,
		*command.NewLoginMetadata(meta.IP, meta.UserAgent, meta.DeviceInfo),
	)

	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.authBus.TwoFactorLogin(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}
}

func (ctrl *TwoFAAuthController) Disable2FA(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	cmd := command.NewDisable2FACommand(userID.Value())
	result := ctrl.twoFaBus.Disable2FA(c.Request.Context(), *cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *TwoFAAuthController) Enable2FA(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	method := c.Query("method")
	if method == "" {
		response.BadRequest(c, errors.New("method query parameter is required"))
		return
	}

	cmd, err := command.NewEnable2FACommand(userID.Value(), method)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.twoFaBus.Enable2FA(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
