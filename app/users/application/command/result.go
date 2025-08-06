package userCommand

type CommandResult struct {
	IsSuccess bool
	Id        string
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

func SuccesResult(id, message string) CommandResult {
	return CommandResult{
		IsSuccess: true,
		Id:        id,
		Message:   message,
		Error:     nil,
	}
}
