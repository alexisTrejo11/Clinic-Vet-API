package session

import (
	"time"
)

type Session struct {
	Id           string
	UserId       string
	RefreshToken string
	DeviceInfo   string
	UserAgent    string
	IpAddress    string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
