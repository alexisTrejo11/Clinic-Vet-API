package auth 

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
)

type Session struct {
	base.Entity
	id           string
	userID       string
	refreshToken string
	deviceInfo   string
	userAgent    string
	ipAddress    string
	expiresAt    time.Time
}
