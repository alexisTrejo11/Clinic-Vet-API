package infraerr

import "net/http"

type ExternalServiceError struct {
	BaseInfrastructureError
	Service string `json:"service"`
}

func NewExternalServiceError(service, message string) *ExternalServiceError {
	return &ExternalServiceError{
		BaseInfrastructureError: BaseInfrastructureError{
			Code:       "EXTERNAL_SERVICE_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusBadGateway,
			Details: map[string]string{
				"service": service,
			},
		},
		Service: service,
	}
}
