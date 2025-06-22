package domain_error

type BaseDomainError struct {
	Code    string                 `json:"code"`
	Type    string                 `json:"type"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"details,omitempty"`
}

func (e BaseDomainError) Error() string {
	return e.Message
}

func (e BaseDomainError) ErrorCode() string {
	return e.Code
}

func (e BaseDomainError) ErrorType() string {
	return e.Type
}

func (e BaseDomainError) Details() map[string]interface{} {
	return e.Data
}
