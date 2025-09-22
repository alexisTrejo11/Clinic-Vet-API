package result

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/auth"
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
	session *SessionResponse
}

func AuthFailure(message string, err error) AuthCommandResult {
	result := &AuthCommandResult{
		CommandResult: *cqrs.FailureResult(message, err),
		session:       nil,
	}
	return *result
}

func AuthSuccess(session *SessionResponse, userid, message string) AuthCommandResult {
	result := &AuthCommandResult{
		CommandResult: *cqrs.SuccessResult(userid, message),
		session:       session,
	}
	return *result
}

func (r *AuthCommandResult) Session() *SessionResponse {
	return r.session
}

func GetSessionResponse(entity auth.Session, access string) *SessionResponse {
	return &SessionResponse{
		UserID:       entity.UserID,
		AccessToken:  access,
		RefreshToken: entity.RefreshToken,
		ExpiresAt:    entity.ExpiresAt,
		CreatedAt:    entity.CreatedAt,
	}
}
