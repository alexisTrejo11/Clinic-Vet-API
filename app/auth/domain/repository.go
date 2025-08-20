package session

import (
	"context"
)

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	GetById(ctx context.Context, sessionId string) (Session, error)
	GetByUserAndId(ctx context.Context, userId, token string) (Session, error)
	GetByUserId(ctx context.Context, userID string) ([]Session, error)
	DeleteUserSession(ctx context.Context, userId, sessionId string) error
	DeleteAllUserSessions(ctx context.Context, userID string) error
}
