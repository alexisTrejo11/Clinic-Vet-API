package response

import (
	"net/http"

	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data any) {
	response := APIResponse{}
	response.SuccessRequest(data)

	ctx.JSON(200, response)
}

func SuccessWithMeta(ctx *gin.Context, data any, meta any) {
	response := APIResponse{}
	response.SuccessWithMeta(data, meta)

	ctx.JSON(200, response)
}

func Created(ctx *gin.Context, data any) {
	response := APIResponse{}
	response.SuccessRequest(data)

	ctx.JSON(201, response)
}

func NoContent(ctx *gin.Context) {
	response := APIResponse{}
	response.SuccessRequest(nil)
	ctx.JSON(204, response)
}

func NotFound(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)
	ctx.JSON(404, errorResponse)
}

func Conflict(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)

	ctx.JSON(409, errorResponse)
}

func BadRequest(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)

	ctx.JSON(400, errorResponse)
}

func Unauthorized(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)

	ctx.JSON(401, errorResponse)
}

func Forbidden(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)

	ctx.JSON(403, errorResponse)
}

func ServerError(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)
	ctx.JSON(500, errorResponse)
}

func ApplicationError(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)
	httpStatusCode := http.StatusInternalServerError

	switch e := err.(type) {
	case interface{ HTTPStatus() int }:
		ctx.JSON(e.HTTPStatus(), e)
		return
	}

	switch e := err.(type) {
	case domainerr.BaseDomainError:
		ctx.JSON(e.HTTPStatus(), e)
		return
	case apperror.BaseApplicationError:
		ctx.JSON(e.HTTPStatus(), e)
		return
	case infraerr.BaseInfrastructureError:
		ctx.JSON(e.HTTPStatus(), e)
		return
	}

	ctx.JSON(httpStatusCode, errorResponse)
}
