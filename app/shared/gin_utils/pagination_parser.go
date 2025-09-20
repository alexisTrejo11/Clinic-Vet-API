package ginutils

import (
	htttpError "clinic-vet-api/app/shared/error/infrastructure/http"
	"clinic-vet-api/app/shared/page"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ShouldBindPageParams(requestPageParams *page.PageInput, ctx *gin.Context, validator *validator.Validate) error {
	if err := ctx.ShouldBindQuery(requestPageParams); err != nil {
		return htttpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery)
	}

	if err := validator.Struct(requestPageParams); err != nil {
		return htttpError.InvalidDataError(err)
	}

	requestPageParams.SetDefaultsFieldsIfEmpty()

	return nil
}
