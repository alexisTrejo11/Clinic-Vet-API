// Package auth contains the authentication related entities like Session and TwoFaAuth
package auth

import (
	"time"
)

type Session struct {
	ID           string
	UserID       string
	RefreshToken string
	DeviceInfo   string
	UserAgent    string
	IPAddress    string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
