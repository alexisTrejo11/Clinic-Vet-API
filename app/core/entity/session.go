package entity

import "time"

type Session struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	DeviceInfo   string    `json:"device_info"`
	UserAgent    string    `json:"user_agent"`
	IpAddress    string    `json:"ip_address"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type TwoFactorAuth struct {
	IsEnabled bool
	Secret    string
}
