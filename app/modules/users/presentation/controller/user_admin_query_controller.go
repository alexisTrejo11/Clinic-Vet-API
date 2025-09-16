package controller

import (
	"errors"

	"clinic-vet-api/app/modules/users/application/usecase/query"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
)

// GetUserByID retrieves a user by their ID.
// @Summary Get a user by ID
// @Description Retrieves a single user record by their unique ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse{data=object} "User found"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Router /v1/users/{id} [get]
func (ctrl *UserAdminController) GetUserByID(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	getUserByIDQuery := query.NewFindUserByIDQuery(c.Request.Context(), userID, false)
	userPage, err := ctrl.queryBus.Execute(getUserByIDQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, userPage, "User")
}

func (ctrl *UserAdminController) SearchUsers(c *gin.Context) {
}

func (ctrl *UserAdminController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		response.BadRequest(c, httpError.RequestURLParamError(errors.New("email is required"), "email", ""))
		return
	}

	getUserByEmailQuery, err := query.NewFindUserByEmailQuery(c.Request.Context(), email, false)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}
	user, err := ctrl.queryBus.Execute(getUserByEmailQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, user, "User")
}

func (c *UserAdminController) GetUserByPhone(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if phone == "" {
		response.BadRequest(ctx, httpError.RequestURLParamError(errors.New("phone is required"), "phone", ""))
		return
	}

	getUserByPhoneQuery, err := query.NewFindUserByPhoneQuery(ctx.Request.Context(), phone, false)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}
	user, err := c.queryBus.Execute(getUserByPhoneQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Found(ctx, user, "User")
}
