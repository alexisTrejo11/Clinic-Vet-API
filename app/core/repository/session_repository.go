package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type SessionRepository interface {
	Create(ctx context.Context, session *auth.Session) error
	GetByID(ctx context.Context, sessionID string) (auth.Session, error)
	GetByUserAndID(ctx context.Context, userID valueobject.UserID, token string) (auth.Session, error)
	GetByUserID(ctx context.Context, userID valueobject.UserID) ([]auth.Session, error)
	DeleteUserSession(ctx context.Context, userID valueobject.UserID, sessionID string) error
	DeleteAllUserSessions(ctx context.Context, userID valueobject.UserID) error
}
