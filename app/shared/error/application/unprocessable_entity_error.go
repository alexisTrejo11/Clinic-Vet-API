package apperror

import (
	"fmt"
	"net/http"
	"strings"
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
			Data: map[string]string{
				"field": field,
				"value": value,
			},
			StatusCode: 422,
		},
	}
}

func MappingError(errorMessage []string, from, to, entity string) *FieldDataError {
	return &FieldDataError{
		BaseApplicationError: BaseApplicationError{
			Code:    "FIELD_DATA_ERROR",
			Type:    "mapping",
			Message: "error courred while mapping " + from,
			Data: map[string]string{
				"errorsMessages": strings.Join(errorMessage, ","),
				"from":           from,
				"to":             to,
				"entity":         entity,
			},
			StatusCode: http.StatusUnprocessableEntity,
		},
	}
}

func BuissnesRuleError(entity, ruleViolated string) *BaseApplicationError {
	return &BaseApplicationError{
		Code:    "BUISNESS_RULE_VIOLATED",
		Type:    "Buisness Logic",
		Message: fmt.Sprintf("Buisness logic fail for %s . RuleViolated: %v", entity, ruleViolated),
		Data: map[string]string{
			"entity":  entity,
			"details": ruleViolated,
		},
		StatusCode: http.StatusUnprocessableEntity,
	}
}
