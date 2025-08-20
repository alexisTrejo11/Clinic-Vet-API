package authCmd

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type SessionResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserId       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthCommandResult struct {
	shared.CommandResult
	Session SessionResponse
}
