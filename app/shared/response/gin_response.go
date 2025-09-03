package response

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data string) {
	response := APIResponse{}
	response.SuccessRequest(data)

	ctx.JSON(200, response)
}

func SuccessWithMeta(ctx *gin.Context, data string, meta any) {
	response := APIResponse{}
	response.SuccessWithMeta(data, meta)

	ctx.JSON(200, response)
}

func Created(ctx *gin.Context, data string) {
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

	fmt.Println(err.Error())
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

	if errors.Is(err, sql.ErrNoRows) {
		httpStatusCode = http.StatusNotFound
	}

	ctx.JSON(httpStatusCode, errorResponse)
}

func RequestURLParamError(ctx *gin.Context, err error, field string, value string) {
	bodyErro := apperror.BaseApplicationError{
		Code:       "INVALID_URL_PARAM",
		Type:       "ROUTING",
		Message:    err.Error(),
		Data:       map[string]string{},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func RequestURLQueryError(ctx *gin.Context, err error) {
	bodyErro := apperror.BaseApplicationError{
		Code:       "INVALID_URL_PARAM",
		Type:       "ROUTING",
		Message:    err.Error(),
		Data:       map[string]string{},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func InvalidDateFormatError(ctx *gin.Context, field, format string) {
	bodyErro := apperror.BaseApplicationError{
		Code:    "INVALID_DATE_FORMAT",
		Type:    "DATE FORMAT",
		Message: fmt.Sprintf("Invalid date format, expected format: %s", format),
		Data: map[string]string{
			"field":  field,
			"format": format,
		},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func InvalidParseDataError(ctx *gin.Context, field string, value, meesage string) {
	bodyErro := apperror.BaseApplicationError{
		Code:    "INVALID_PARSE_DATA",
		Type:    "DATA PARSING",
		Message: fmt.Sprintf("Invalid data for field '%s': %s", field, value),
		Data: map[string]string{
			"field":   field,
			"value":   value,
			"message": meesage,
		},

		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func RequestBodyDataError(ctx *gin.Context, err error) {
	bodyErro := apperror.BaseApplicationError{
		Code:       "REQUEST_BODY_ERROR",
		Type:       "BODY FORMAT",
		Message:    err.Error(),
		Data:       map[string]string{},
		StatusCode: 400,
	}
	BadRequest(ctx, bodyErro)
}

func GetPaginationParams(ctx *gin.Context) (int, int, error) {
	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

	pageNumber, err := strconv.Atoi(pageParam)
	if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 0, 0, err
	}

	if pageNumber < 1 {
		return 0, 0, errors.New("page number must be greater than 0")
	}

	if limit < 1 {
		return 0, 0, errors.New("limit must be greater than 0")
	}

	return pageNumber, limit, nil
}
