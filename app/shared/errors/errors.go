package custom_error

type DomainError interface {
	error
	ErrorCode() string
	ErrorType() string
	Details() map[string]interface{}
}

type ApplicationError interface {
	error
	ErrorCode() string
	ErrorType() string
	Details() map[string]interface{}
	HTTPStatus() int
}

type InfrastructureError interface {
	error
	ErrorCode() string
	ErrorType() string
	Details() map[string]interface{}
	HTTPStatus() int
}
