package sessionRepo

import (
	"context"

	sessionDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
)

type SessionRepository interface {
	Create(ctx context.Context, session *sessionDomain.Session) error
	FindByRefreshToken(ctx context.Context, sessionId string) (*sessionDomain.Session, error)
	FindByUserAndRefreshToken(ctx context.Context, userId, token string) (*sessionDomain.Session, error)
	FindByUserID(ctx context.Context, userID int) ([]*sessionDomain.Session, error)
	UpdateSession(ctx context.Context, session *sessionDomain.Session) error
	DeleteUserSession(ctx context.Context, userId, sessionId string) error
	DeleteAllUserSessions(ctx context.Context, userID string) error
}
