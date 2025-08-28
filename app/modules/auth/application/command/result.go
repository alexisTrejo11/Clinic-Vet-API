package authCmd

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type SessionResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserID       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthCommandResult struct {
	result  shared.CommandResult
	session SessionResponse
}

func FailureAuthResult(message string, err error) AuthCommandResult {
	result := &AuthCommandResult{
		result:  shared.FailureResult(message, err),
		session: SessionResponse{},
	}

	return *result
}

func SuccessAuthResult(session *SessionResponse, userid, message string) AuthCommandResult {
	result := &AuthCommandResult{
		result:  shared.SuccessResult(userid, message),
		session: *session,
	}

	return *result
}

func (r *AuthCommandResult) IsSuccess() bool {
	return r.result.IsSuccess
}

func (r *AuthCommandResult) Error() error {
	return r.result.Error
}

func (r *AuthCommandResult) Message() string {
	return r.result.Message
}

func (r *AuthCommandResult) Session() SessionResponse {
	return r.Session()
}
