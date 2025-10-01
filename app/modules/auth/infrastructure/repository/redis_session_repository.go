package repositoryimpl

import (
	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var sessionDuration time.Duration = (24 * time.Hour) * 30 // Session expiration duration

type RedisSessionRepository struct {
	redisClient *redis.Client
}

func NewRedisSessionRepository(redisClient *redis.Client) repository.SessionRepository {
	return &RedisSessionRepository{
		redisClient: redisClient,
	}
}

// sessionKey returns the Redis key for a single session.
// Example: "session:jdn23n-k4n2k4-n2k34n-k3m4n2"
func (r *RedisSessionRepository) sessionKey(refresh string) string {
	return fmt.Sprintf("session:%s", refresh)
}

// userSessionsKey returns the Redis key for a set of sessions for a specific userDomain.
// Example: "user_sessions:123"
func (r *RedisSessionRepository) userSessionsKey(userID string) string {
	return fmt.Sprintf("user_sessions:%s", userID)
}

func (r *RedisSessionRepository) Create(ctx context.Context, sess *auth.Session) error {
	sessionJSON, err := json.Marshal(sess)
	if err != nil {
		return fmt.Errorf("failed to marshal session to JSON: %w", err)
	}

	_, err = r.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		// Store the session data with an expiration time
		pipe.Set(ctx, r.sessionKey(sess.RefreshToken), sessionJSON, sessionDuration)

		// Add the session ID to the set of user sessions
		pipe.SAdd(ctx, r.userSessionsKey(sess.UserID), sess.RefreshToken)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to create session in Redis: %w", err)
	}

	return nil
}

func (r *RedisSessionRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (auth.Session, error) {
	sessionJSON, err := r.redisClient.Get(ctx, r.sessionKey(refreshToken)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return auth.Session{}, fmt.Errorf("session not found for Refresh Token: %s", refreshToken[:3])
		}
		return auth.Session{}, fmt.Errorf("failed to get session from Redis: %w", err)
	}

	var sess auth.Session
	if err := json.Unmarshal(sessionJSON, &sess); err != nil {
		return auth.Session{}, fmt.Errorf("failed to unmarshal session JSON: %w", err)
	}

	return sess, nil
}

func (r *RedisSessionRepository) GetByUserAndRefreshToken(ctx context.Context, userID valueobject.UserID, token string) (auth.Session, error) {
	isMember, err := r.redisClient.SIsMember(ctx, r.userSessionsKey(userID.String()), token).Result()
	if err != nil {
		return auth.Session{}, fmt.Errorf("failed to check session membership: %w", err)
	}
	if !isMember {
		return auth.Session{}, fmt.Errorf("session with token %s not found for user %s", token, userID)
	}

	// Then, retrieve the session data itself
	return r.GetByRefreshToken(ctx, token)
}

func (r *RedisSessionRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) ([]auth.Session, error) {
	sessionIDs, err := r.redisClient.SMembers(ctx, r.userSessionsKey(userID.String())).Result()
	if err != nil {
		return []auth.Session{}, fmt.Errorf("failed to get session IDs for user %s: %w", userID, err)
	}

	if len(sessionIDs) == 0 {
		return []auth.Session{}, nil
	}

	sessionKeys := make([]string, len(sessionIDs))
	for i, id := range sessionIDs {
		sessionKeys[i] = r.sessionKey(id)
	}

	sessionData, err := r.redisClient.MGet(ctx, sessionKeys...).Result()
	if err != nil {
		return []auth.Session{}, fmt.Errorf("failed to get sessions from Redis: %w", err)
	}

	sessions := make([]auth.Session, 0, len(sessionData))
	for _, val := range sessionData {
		if val == nil {
			continue // Skip if a session key expired or was deleted
		}

		var sess auth.Session
		if err := json.Unmarshal([]byte(val.(string)), &sess); err != nil {
			return []auth.Session{}, fmt.Errorf("failed to unmarshal session JSON: %w", err)
		}
		sessions = append(sessions, sess)
	}
	return sessions, nil
}

func (r *RedisSessionRepository) DeleteUserSession(ctx context.Context, userID valueobject.UserID, sessionID string) error {
	_, err := r.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, r.sessionKey(sessionID))
		pipe.SRem(ctx, r.userSessionsKey(userID.String()), sessionID)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete session from Redis: %w", err)
	}

	return nil
}

func (r *RedisSessionRepository) DeleteAllUserSessions(ctx context.Context, userID valueobject.UserID) error {
	sessionIDs, err := r.redisClient.SMembers(ctx, r.userSessionsKey(userID.String())).Result()
	if err != nil {
		return fmt.Errorf("failed to get sessions for user: %w", err)
	}

	_, err = r.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, sessionID := range sessionIDs {
			pipe.Del(ctx, r.sessionKey(sessionID))
		}
		pipe.Del(ctx, r.userSessionsKey(userID.String()))
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete all sessions for user %s: %w", userID, err)
	}

	return nil
}
