package cqrs

type CommandResult struct {
	IsSuccess bool
	ID        string
	Message   string
	Error     error
}

func FailureResult(message string, err error) CommandResult {
	return CommandResult{
		IsSuccess: false,
		Message:   message,
		Error:     err,
	}
}

func SuccessResult(id, message string) CommandResult {
	return CommandResult{
		IsSuccess: true,
		ID:        id,
		Message:   message,
		Error:     nil,
	}
}
