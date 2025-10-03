package cqrs

import (
	"encoding/json"
)

type CommandResult struct {
	isSuccess bool
	id        string
	message   string
	error     error
}

func (c CommandResult) IsSuccess() bool {
	return c.isSuccess
}

func (c CommandResult) ID() string {
	return c.id
}

func (c CommandResult) Message() string {
	return c.message
}

func (c CommandResult) Error() error {
	return c.error
}

func (c CommandResult) ToMap() map[string]any {
	result := map[string]any{
		"isSuccess": c.isSuccess,
		"id":        c.id,
		"message":   c.message,
	}
	if c.error != nil {
		result["error"] = c.error.Error()
	}
	return result
}

func (c CommandResult) MarshalJSON() ([]byte, error) {
	type Alias CommandResult
	return json.Marshal(&struct {
		ErrorString string `json:"error,omitempty"`
		*Alias
	}{
		ErrorString: c.getErrorString(),
		Alias:       (*Alias)(&c),
	})
}

func (c CommandResult) getErrorString() string {
	if c.error != nil {
		return c.error.Error()
	}
	return ""
}

func FailureResult(message string, err error) CommandResult {
	return CommandResult{
		isSuccess: false,
		message:   message,
		error:     err,
	}
}

func SuccessCreateResult(id string, message string) CommandResult {
	return CommandResult{
		isSuccess: true,
		id:        id,
		message:   message,
		error:     nil,
	}
}

func SuccessResult(message string) CommandResult {
	return CommandResult{
		isSuccess: true,
		message:   message,
		error:     nil,
	}
}
