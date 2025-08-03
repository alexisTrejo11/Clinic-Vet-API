package authRepository

import (
	sessionDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
)

type SessionRepository interface {
	Create(session *sessionDomain.Session) error
	FindByRefreshToken(token string) (*sessionDomain.Session, error)
	FindByUserID(userID int) ([]*sessionDomain.Session, error)
	UpdateSession(session *sessionDomain.Session) error
	DeleteSession(sessionID int) error
	DeleteAllUserSessions(userID int) error
	DeleteExpiredSessions() error
}
