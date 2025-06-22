package appError

type BaseApplicationError struct {
	Code       string                 `json:"code"`
	Type       string                 `json:"type"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"details,omitempty"`
	StatusCode int                    `json:"-"`
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

func (e BaseApplicationError) Details() map[string]interface{} {
	return e.Data
}

func (e BaseApplicationError) HTTPStatus() int {
	return e.StatusCode
}
