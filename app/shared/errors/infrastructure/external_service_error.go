package infraErr

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
			Data: map[string]interface{}{
				"service": service,
			},
		},
		Service: service,
	}
}
