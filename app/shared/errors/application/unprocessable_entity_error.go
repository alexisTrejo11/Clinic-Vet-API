package apperror

import (
	"fmt"
	"net/http"
)

type FieldDataError struct {
	BaseApplicationError
}

func FieldValidationError(field, value, message string) *FieldDataError {
	return &FieldDataError{
		BaseApplicationError: BaseApplicationError{
			Code:    "FIELD_DATA_ERROR",
			Type:    "mapping",
			Message: message,
			Data: map[string]any{
				"field": field,
				"value": value,
			},
			StatusCode: 422,
		},
	}
}

func MappingError(errorMap map[string]any, from, to string) *FieldDataError {
	return &FieldDataError{
		BaseApplicationError: BaseApplicationError{
			Code:       "FIELD_DATA_ERROR",
			Type:       "mapping",
			Message:    "error courred while mapping " + from + " --> " + to,
			Data:       errorMap,
			StatusCode: http.StatusUnprocessableEntity,
		},
	}
}

func BuissnesRuleError(entity, ruleViolated string) *BaseApplicationError {
	return &BaseApplicationError{
		Code:    "BUISNESS_RULE_VIOLATED",
		Type:    "Buisness Logic",
		Message: fmt.Sprintf("Buisness logic fail for %s . RuleViolated: %v", entity, ruleViolated),
		Data: map[string]any{
			"entity":  entity,
			"details": ruleViolated,
		},
		StatusCode: http.StatusUnprocessableEntity,
	}
}
