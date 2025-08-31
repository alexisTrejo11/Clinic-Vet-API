package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
	GetByID(ctx context.Context, sessionID string) (entity.Session, error)
	GetByUserAndID(ctx context.Context, userID, token string) (entity.Session, error)
	GetByUserID(ctx context.Context, userID string) ([]entity.Session, error)
	DeleteUserSession(ctx context.Context, userID, sessionID string) error
	DeleteAllUserSessions(ctx context.Context, userID string) error
}
