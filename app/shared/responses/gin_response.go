package apiResponse

import (
	"fmt"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/gin-gonic/gin"
)

func Ok(ctx *gin.Context, data interface{}) {
	response := APIResponse{}
	response.SuccessRequest(data)

	ctx.JSON(200, response)
}

func Created(ctx *gin.Context, data interface{}) {
	response := APIResponse{}
	response.SuccessRequest(data)

	ctx.JSON(201, response)
}

func NoContent(ctx *gin.Context) {
	response := APIResponse{}
	response.SuccessRequest(nil)
	ctx.JSON(204, response)
}

func NoFound(ctx *gin.Context, err error) {
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

	fmt.Println(err.Error())
	ctx.JSON(400, errorResponse)
}

func Unauthorized(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)

	ctx.JSON(401, errorResponse)
}

func Forbbbiden(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)

	ctx.JSON(403, errorResponse)
}

func ServerError(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)
	ctx.JSON(500, errorResponse)
}

func AppError(ctx *gin.Context, err error) {
	response := APIResponse{}
	errorResponse := response.ErrorRequest(err)
	httpStatusCode := 500

	type httpStatusCoder interface {
		HTTPStatus() int
	}

	if e, ok := err.(httpStatusCoder); ok {
		httpStatusCode = e.HTTPStatus()
	}

	ctx.JSON(httpStatusCode, errorResponse)
}

func RequestURLParamError(ctx *gin.Context, err error, field string, value string) {
	bodyErro := appError.BaseApplicationError{
		Code:       "INVALID_URL_PARAM",
		Type:       "ROUTING",
		Message:    err.Error(),
		Data:       map[string]interface{}{},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func RequestURLQueryError(ctx *gin.Context, err error) {
	bodyErro := appError.BaseApplicationError{
		Code:       "INVALID_URL_PARAM",
		Type:       "ROUTING",
		Message:    err.Error(),
		Data:       map[string]interface{}{},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func RequestBodyDataError(ctx *gin.Context, err error) {
	bodyErro := appError.BaseApplicationError{
		Code:       "REQUEST_BODY_ERROR",
		Type:       "BODY FORMAT",
		Message:    err.Error(),
		Data:       map[string]interface{}{},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}
