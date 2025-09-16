package ginutils

import (
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BindAndValidate binds the request body to the provided object and validates it using the provided validation function.
func BindAndValidateBody(c *gin.Context, obj any, validate *validator.Validate) error {
	if err := c.ShouldBindBodyWithJSON(obj); err != nil {
		return httpError.RequestBodyDataError(err)
	}

	if err := validate.Struct(obj); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return httpError.InvalidDataError(validationErrors)
	}

	return nil
}

func BindAndValidateQuery(c *gin.Context, obj any, validate *validator.Validate) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return httpError.InvalidDataError(err)
	}

	if err := validate.Struct(obj); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return httpError.InvalidDataError(validationErrors)
	}

	return nil
}
