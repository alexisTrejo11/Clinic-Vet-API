package sessionRepository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sessionDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	"github.com/redis/go-redis/v9"
)

var sessionDuration time.Duration = (24 * time.Hour) * 30 // Session expiration duration

type RedisSessionRepository struct {
	redisClient *redis.Client
}

func NewRedisSessionRepository(redisClient *redis.Client) sessionDomain.SessionRepository {
	return &RedisSessionRepository{
		redisClient: redisClient,
	}
}

// sessionKey returns the Redis key for a single session.
// Example: "session:jdn23n-k4n2k4-n2k34n-k3m4n2"
func (r *RedisSessionRepository) sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

// userSessionsKey returns the Redis key for a set of sessions for a specific userDomain.
// Example: "user_sessions:123"
func (r *RedisSessionRepository) userSessionsKey(userId string) string {
	return fmt.Sprintf("user_sessions:%s", userId)
}

func (r *RedisSessionRepository) Create(ctx context.Context, sess *sessionDomain.Session) error {
	sessionJSON, err := json.Marshal(sess)
	if err != nil {
		return fmt.Errorf("failed to marshal session to JSON: %w", err)
	}

	_, err = r.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		// Store the session data with an expiration time
		pipe.Set(ctx, r.sessionKey(sess.Id), sessionJSON, sessionDuration)

		// Add the session ID to the set of user sessions
		pipe.SAdd(ctx, r.userSessionsKey(sess.UserId), sess.Id)
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to create session in Redis: %w", err)
	}

	return nil
}
func (r *RedisSessionRepository) GetById(ctx context.Context, sessionId string) (sessionDomain.Session, error) {
	sessionJSON, err := r.redisClient.Get(ctx, r.sessionKey(sessionId)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return sessionDomain.Session{}, fmt.Errorf("session not found for ID: %s", sessionId)
		}
		return sessionDomain.Session{}, fmt.Errorf("failed to get session from Redis: %w", err)
	}

	var sess sessionDomain.Session
	if err := json.Unmarshal(sessionJSON, &sess); err != nil {
		return sessionDomain.Session{}, fmt.Errorf("failed to unmarshal session JSON: %w", err)
	}

	return sess, nil
}

func (r *RedisSessionRepository) GetByUserAndId(ctx context.Context, userId, token string) (sessionDomain.Session, error) {
	isMember, err := r.redisClient.SIsMember(ctx, r.userSessionsKey(userId), token).Result()
	if err != nil {
		return sessionDomain.Session{}, fmt.Errorf("failed to check session membership: %w", err)
	}
	if !isMember {
		return sessionDomain.Session{}, fmt.Errorf("session with token %s not found for user %s", token, userId)
	}

	// Then, retrieve the session data itself
	return r.GetById(ctx, token)

}

func (r *RedisSessionRepository) GetByUserId(ctx context.Context, userId string) ([]sessionDomain.Session, error) {
	sessionIds, err := r.redisClient.SMembers(ctx, r.userSessionsKey(userId)).Result()
	if err != nil {
		return []sessionDomain.Session{}, fmt.Errorf("failed to get session IDs for user %s: %w", userId, err)
	}

	if len(sessionIds) == 0 {
		return []sessionDomain.Session{}, nil
	}

	sessionKeys := make([]string, len(sessionIds))
	for i, id := range sessionIds {
		sessionKeys[i] = r.sessionKey(id)
	}

	sessionData, err := r.redisClient.MGet(ctx, sessionKeys...).Result()
	if err != nil {
		return []sessionDomain.Session{}, fmt.Errorf("failed to get sessions from Redis: %w", err)
	}

	sessions := make([]sessionDomain.Session, 0, len(sessionData))
	for _, val := range sessionData {
		if val == nil {
			continue // Skip if a session key expired or was deleted
		}

		var sess sessionDomain.Session
		if err := json.Unmarshal([]byte(val.(string)), &sess); err != nil {
			return []sessionDomain.Session{}, fmt.Errorf("failed to unmarshal session JSON: %w", err)
		}
		sessions = append(sessions, sess)
	}
	return sessions, nil
}

func (r *RedisSessionRepository) DeleteUserSession(ctx context.Context, userId, sessionId string) error {
	_, err := r.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, r.sessionKey(sessionId))
		pipe.SRem(ctx, r.userSessionsKey(userId), sessionId)
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to delete session from Redis: %w", err)
	}

	return nil
}

func (r *RedisSessionRepository) DeleteAllUserSessions(ctx context.Context, userID string) error {
	sessionIds, err := r.redisClient.SMembers(ctx, r.userSessionsKey(userID)).Result()
	if err != nil {
		return fmt.Errorf("failed to get sessions for user: %w", err)
	}

	_, err = r.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, sessionId := range sessionIds {
			pipe.Del(ctx, r.sessionKey(sessionId))
		}
		pipe.Del(ctx, r.userSessionsKey(userID))
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to delete all sessions for user %s: %w", userID, err)
	}

	return nil
}
