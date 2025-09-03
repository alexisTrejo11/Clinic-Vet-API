package repositoryimpl

import (
	"net/http"

	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
)

func AppointmentDBError(message string) error {
	return &infraerr.DatabaseError{
		BaseInfrastructureError: infraerr.BaseInfrastructureError{
			Code:       "APPOINTMENT_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Details: map[string]string{
				"operation": "appointment database operation",
				"table":     "appointments",
			},
		},
		Operation: "appointment database operation",
	}
}

func AppointmentInsertDBErr(message string) error {
	return &infraerr.DatabaseError{
		BaseInfrastructureError: infraerr.BaseInfrastructureError{
			Code:       "APPOINTMENT_INSERT_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Details: map[string]string{
				"operation": "insert appointment to database",
				"table":     "appointments",
			},
		},
		Operation: "insert appointment to database",
	}
}

func AppointmentUpdateDBErr(message string) error {
	return &infraerr.DatabaseError{
		BaseInfrastructureError: infraerr.BaseInfrastructureError{
			Code:       "APPOINTMENT_UPDATE_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Details: map[string]string{
				"operation": "update appointment in database",
				"table":     "appointments",
			},
		},
		Operation: "update appointment in database",
	}
}

func AppointmentDeleteDBErr(message string) error {
	return &infraerr.DatabaseError{
		BaseInfrastructureError: infraerr.BaseInfrastructureError{
			Code:       "APPOINTMENT_DELETE_DB_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Details: map[string]string{
				"operation": "delete appointment from database",
				"table":     "appointments",
			},
		},
		Operation: "delete appointment from database",
	}
}

func AppointmentNotFoundErr(param, value string) error {
	return &infraerr.DatabaseError{
		BaseInfrastructureError: infraerr.BaseInfrastructureError{
			Code:       "APPOINTMENT_NOT_FOUND",
			Type:       "infrastructure",
			Message:    "Appointment" + " with " + param + " '" + value + "' not found",
			StatusCode: http.StatusNotFound,
			Details: map[string]string{
				"operation": "find appointment in database",
				"param":     param,
				"value":     value,
				"table":     "appointments",
			},
		},
		Operation: "find appointment in database",
	}
}
