package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/account/auth/local/application/command"
	"clinic-vet-api/app/modules/account/user/infrastructure/bus"
	"clinic-vet-api/app/modules/account/user/presentation/dto"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/error/infrastructure/http"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserUpdateController struct {
	validator *validator.Validate
	bus       *bus.UserBus
}

func NewUserUpdateController(validator *validator.Validate, bus *bus.UserBus) *UserUpdateController {
	return &UserUpdateController{
		validator: validator,
		bus:       bus,
	}
}

func (ctrl *UserUpdateController) SendResetPasswordEmail(c *gin.Context) {
	var req dto.RequestEmailRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	command, err := req.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.RequestResetPassword(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *UserUpdateController) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	command, err := req.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.ResetPassword(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *UserUpdateController) UpdatePassword(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.BadRequest(c, autherror.UnauthorizedCTXError())
		return
	}

	var req dto.UpdatePasswordRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	cmd, err := req.ToCommand(id.Value())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.UpdatePassword(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *UserUpdateController) UpdateEmail(c *gin.Context) {

}

func (ctrl *UserUpdateController) UpdatePhoneNumber(c *gin.Context) {

}

func (ctrl *UserUpdateController) ActivateAccount(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil && userID == 0 {
		response.BadRequest(c, http.MissingRequiredFieldError("user id is required"))
		return
	}

	code := c.Query("code")
	if code == "" {
		response.BadRequest(c, http.MissingRequiredFieldError("activation code is required"))
		return
	}

	command, err := command.NewActivateAccountCommand(code, userID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.ActivateAccount(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
