package infra_error

type BaseInfrastructureError struct {
	Code       string                 `json:"code"`
	Type       string                 `json:"type"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"details,omitempty"`
	StatusCode int                    `json:"-"`
}

func (e BaseInfrastructureError) Error() string {
	return e.Message
}

func (e BaseInfrastructureError) ErrorCode() string {
	return e.Code
}

func (e BaseInfrastructureError) ErrorType() string {
	return e.Type
}

func (e BaseInfrastructureError) Details() map[string]interface{} {
	return e.Data
}

func (e BaseInfrastructureError) HTTPStatus() int {
	return e.StatusCode
}
