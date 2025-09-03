package cqrs

import (
	"encoding/json"
)

type CommandResult struct {
	IsSuccess bool   `json:"isSuccess"`
	ID        string `json:"id,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     error  `json:"error,omitempty"`
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
	if c.Error != nil {
		return c.Error.Error()
	}
	return ""
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

func (c CommandResult) ToMap() map[string]any {
	result := map[string]any{
		"isSuccess": c.IsSuccess,
		"id":        c.ID,
		"message":   c.Message,
	}

	if c.Error != nil {
		result["error"] = c.Error.Error()
	}

	return result
}
