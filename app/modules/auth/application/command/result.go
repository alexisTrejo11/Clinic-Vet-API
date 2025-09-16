package command

import (
	"time"

	"clinic-vet-api/app/shared/cqrs"
)

type SessionResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserID       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthCommandResult struct {
	cqrs.CommandResult
	session SessionResponse
}

func FailureAuthResult(message string, err error) AuthCommandResult {
	result := &AuthCommandResult{
		CommandResult: *cqrs.FailureResult(message, err),
		session:       SessionResponse{},
	}
	return *result
}

func SuccessAuthResult(session *SessionResponse, userid, message string) AuthCommandResult {
	result := &AuthCommandResult{
		CommandResult: *cqrs.SuccessResult(userid, message),
		session:       *session,
	}
	return *result
}

func (r *AuthCommandResult) Session() SessionResponse {
	return r.session
}
