package sqlcrepository

import (
	"net/http"

	infraErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/infrastructure"
)

func AppointmentDBError(message string) error {
	return &infraErr.DatabaseError{
		BaseInfrastructureError: infraErr.BaseInfrastructureError{
			Code:       "APPOINTMENT_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Data: map[string]interface{}{
				"operation": "appointment database operation",
				"table":     "appointments",
			},
		},
		Operation: "appointment database operation",
	}
}

func AppointmentInsertDBErr(message string) error {
	return &infraErr.DatabaseError{
		BaseInfrastructureError: infraErr.BaseInfrastructureError{
			Code:       "APPOINTMENT_INSERT_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Data: map[string]interface{}{
				"operation": "insert appointment to database",
				"table":     "appointments",
			},
		},
		Operation: "insert appointment to database",
	}
}

func AppointmentUpdateDBErr(message string) error {
	return &infraErr.DatabaseError{
		BaseInfrastructureError: infraErr.BaseInfrastructureError{
			Code:       "APPOINTMENT_UPDATE_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Data: map[string]interface{}{
				"operation": "update appointment in database",
				"table":     "appointments",
			},
		},
		Operation: "update appointment in database",
	}
}

func AppointmentDeleteDBErr(message string) error {
	return &infraErr.DatabaseError{
		BaseInfrastructureError: infraErr.BaseInfrastructureError{
			Code:       "APPOINTMENT_DELETE_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Data: map[string]interface{}{
				"operation": "delete appointment from database",
				"table":     "appointments",
			},
		},
		Operation: "delete appointment from database",
	}
}

func AppointmentNotFoundErr(param, value string) error {
	return &infraErr.DatabaseError{
		BaseInfrastructureError: infraErr.BaseInfrastructureError{
			Code:       "APPOINTMENT_NOT_FOUND",
			Type:       "infrastructure",
			Message:    "Appointment" + " with " + param + " '" + value + "' not found",
			StatusCode: http.StatusNotFound,
			Data: map[string]interface{}{
				"operation": "find appointment in database",
				"param":     param,
				"value":     value,
				"table":     "appointments",
			},
		},
		Operation: "find appointment in database",
	}
}
