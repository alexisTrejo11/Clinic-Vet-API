package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
)

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
	GetById(ctx context.Context, sessionId string) (entity.Session, error)
	GetByUserAndId(ctx context.Context, userId, token string) (entity.Session, error)
	GetByUserId(ctx context.Context, userID string) ([]entity.Session, error)
	DeleteUserSession(ctx context.Context, userId, sessionId string) error
	DeleteAllUserSessions(ctx context.Context, userID string) error
}
