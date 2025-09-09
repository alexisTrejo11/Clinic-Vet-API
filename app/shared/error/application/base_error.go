// Package apperror contains all the common app errors to be used on application layer
package apperror

type BaseApplicationError struct {
	Code       string            `json:"code"`
	Type       string            `json:"type"`
	Message    string            `json:"message"`
	Details    map[string]string `json:"details,omitempty"`
	StatusCode int               `json:"-"`
}

func (e BaseApplicationError) HTTPStatusCode() int {
	if e.StatusCode == 0 {
		return 500
	}
	return e.StatusCode
}

func (e BaseApplicationError) Error() string {
	return e.Message
}

func (e BaseApplicationError) ErrorCode() string {
	return e.Code
}

func (e BaseApplicationError) ErrorType() string {
	return e.Type
}

func (e BaseApplicationError) DetailMap() map[string]string {
	return e.Details
}

func (e BaseApplicationError) HTTPStatus() int {
	return e.StatusCode
}
