// Package infraerr contains errors for infrastructure errors like database conflicts, external services failures
package infraerr

type BaseInfrastructureError struct {
	Code       string            `json:"code"`
	Type       string            `json:"type"`
	Message    string            `json:"message"`
	Details    map[string]string `json:"details,omitempty"`
	StatusCode int               `json:"-"`
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

func (e BaseInfrastructureError) DetailMap() map[string]string {
	return e.Details
}

func (e BaseInfrastructureError) HTTPStatus() int {
	return e.StatusCode
}
