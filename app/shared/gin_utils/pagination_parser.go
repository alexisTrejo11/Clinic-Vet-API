package ginutils

import (
	htttpError "clinic-vet-api/app/shared/error/infrastructure/http"
	"clinic-vet-api/app/shared/page"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ShouldBindPageParams(requestPageParams *page.PaginationRequest, ctx *gin.Context, validator *validator.Validate) error {
	if err := ctx.ShouldBindQuery(requestPageParams); err != nil {
		return htttpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery)
	}

	*requestPageParams = requestPageParams.WithDefaults()

	if err := validator.Struct(requestPageParams); err != nil {
		return htttpError.InvalidDataError(err)
	}

	return nil
}
