package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

// RedisClientInfo represents client info stored in Redis
type RedisClientInfo struct {
	Requests  []time.Time `json:"requests"`
	LastReset time.Time   `json:"last_reset"`
}

// NewRedisStorage creates a new Redis storage instance
func NewRedisStorage(redisURL string) (*RedisStorage, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	// Test connection
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisStorage{
		client: client,
		ctx:    ctx,
	}, nil
}

func (r *RedisStorage) Get(key string) (*ClientInfo, bool) {
	val, err := r.client.Get(r.ctx, "ratelimit:"+key).Result()
	if err == redis.Nil {
		return nil, false
	}
	if err != nil {
		// Log error and return false to create new entry
		fmt.Printf("Redis GET error: %v\n", err)
		return nil, false
	}

	var redisInfo RedisClientInfo
	if err := json.Unmarshal([]byte(val), &redisInfo); err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		return nil, false
	}

	clientInfo := &ClientInfo{
		Requests:  redisInfo.Requests,
		LastReset: redisInfo.LastReset,
	}

	return clientInfo, true
}

func (r *RedisStorage) Set(key string, info *ClientInfo) {
	redisInfo := RedisClientInfo{
		Requests:  info.Requests,
		LastReset: info.LastReset,
	}

	data, err := json.Marshal(redisInfo)
	if err != nil {
		fmt.Printf("JSON marshal error: %v\n", err)
		return
	}

	// Set with TTL of 1 hour
	err = r.client.Set(r.ctx, "ratelimit:"+key, data, time.Hour).Err()
	if err != nil {
		fmt.Printf("Redis SET error: %v\n", err)
	}
}

func (r *RedisStorage) Delete(key string) {
	err := r.client.Del(r.ctx, "ratelimit:"+key).Err()
	if err != nil {
		fmt.Printf("Redis DEL error: %v\n", err)
	}
}

func (r *RedisStorage) Cleanup() {
	// Redis handles TTL automatically, but we can implement custom cleanup if needed
	keys, err := r.client.Keys(r.ctx, "ratelimit:*").Result()
	if err != nil {
		fmt.Printf("Redis KEYS error: %v\n", err)
		return
	}

	now := time.Now()
	for _, key := range keys {
		val, err := r.client.Get(r.ctx, key).Result()
		if err != nil {
			continue
		}

		var redisInfo RedisClientInfo
		if err := json.Unmarshal([]byte(val), &redisInfo); err != nil {
			continue
		}

		if now.Sub(redisInfo.LastReset) > time.Hour {
			r.client.Del(r.ctx, key)
		}
	}
}

// LuaScript for atomic rate limit check and increment
const rateLimitLuaScript = `
local key = KEYS[1]
local window = tonumber(ARGV[1])
local limit = tonumber(ARGV[2])
local current_time = tonumber(ARGV[3])

-- Get current data
local current = redis.call('HMGET', key, 'count', 'window_start')
local count = tonumber(current[1]) or 0
local window_start = tonumber(current[2]) or current_time

-- Check if we need to reset the window
if current_time - window_start >= window then
    count = 0
    window_start = current_time
end

-- Check if limit exceeded
if count >= limit then
    return {0, count, window_start + window}
end

-- Increment count
count = count + 1
redis.call('HMSET', key, 'count', count, 'window_start', window_start)
redis.call('EXPIRE', key, window)

return {1, count, window_start + window}
`

// AtomicRedisStorage uses Lua scripts for atomic operations
type AtomicRedisStorage struct {
	client    *redis.Client
	ctx       context.Context
	luaScript *redis.Script
}

// NewAtomicRedisStorage creates a Redis storage with atomic operations
func NewAtomicRedisStorage(redisURL string) (*AtomicRedisStorage, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &AtomicRedisStorage{
		client:    client,
		ctx:       ctx,
		luaScript: redis.NewScript(rateLimitLuaScript),
	}, nil
}

// CheckAndIncrement atomically checks rate limit and increments counter
func (a *AtomicRedisStorage) CheckAndIncrement(key string, windowSeconds, limit int) (allowed bool, count int, resetTime time.Time) {
	currentTime := time.Now().Unix()

	result, err := a.luaScript.Run(
		a.ctx,
		a.client,
		[]string{"ratelimit:" + key},
		windowSeconds,
		limit,
		currentTime,
	).Result()

	if err != nil {
		fmt.Printf("Lua script error: %v\n", err)
		return true, 0, time.Now().Add(time.Duration(windowSeconds) * time.Second)
	}

	values := result.([]interface{})
	allowed = values[0].(int64) == 1
	count = int(values[1].(int64))
	resetTime = time.Unix(values[2].(int64), 0)

	return allowed, count, resetTime
}

// AtomicRateLimiter creates a rate limiter using atomic Redis operations
func AtomicRateLimiter(storage *AtomicRedisStorage, requestsPerSecond int, windowDuration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		windowSeconds := int(windowDuration.Seconds())
		limit := int(float64(requestsPerSecond) * windowDuration.Seconds())

		allowed, count, resetTime := storage.CheckAndIncrement(key, windowSeconds, limit)

		// Set headers
		c.Header("X-RateLimit-Limit", strconv.Itoa(limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(limit-count))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(resetTime.Unix(), 10))

		if !allowed {
			c.Header("Retry-After", strconv.FormatInt(int64(time.Until(resetTime).Seconds()), 10))
			c.JSON(429, gin.H{
				"error":       "Rate limit exceeded",
				"retry_after": time.Until(resetTime).Seconds(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Example with Redis configuration
func SetupRedisRateLimiter() gin.HandlerFunc {
	// Initialize Redis storage
	redisStorage, err := NewAtomicRedisStorage("redis://localhost:6379/0")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return AtomicRateLimiter(redisStorage, 100, time.Minute)
}

// Multi-tier rate limiting with Redis
func MultiTierRedisRateLimiter() gin.HandlerFunc {
	storage, err := NewAtomicRedisStorage("redis://localhost:6379/0")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()

		// Tier 1: Per second limit
		allowed, _, _ := storage.CheckAndIncrement(
			fmt.Sprintf("%s:second:%d", clientIP, now.Unix()),
			1,  // 1 second window
			10, // 10 requests per second
		)

		if !allowed {
			c.Header("X-RateLimit-Type", "per-second")
			c.Header("Retry-After", "1")
			c.JSON(429, gin.H{"error": "Rate limit exceeded (per second)"})
			c.Abort()
			return
		}

		// Tier 2: Per minute limit
		allowed, count, resetTime := storage.CheckAndIncrement(
			fmt.Sprintf("%s:minute:%d", clientIP, now.Unix()/60),
			60,  // 1 minute window
			100, // 100 requests per minute
		)

		c.Header("X-RateLimit-Limit", "100")
		c.Header("X-RateLimit-Remaining", strconv.Itoa(100-count))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(resetTime.Unix(), 10))

		if !allowed {
			c.Header("X-RateLimit-Type", "per-minute")
			c.Header("Retry-After", strconv.FormatInt(int64(resetTime.Sub(now).Seconds()), 10))
			c.JSON(429, gin.H{"error": "Rate limit exceeded (per minute)"})
			c.Abort()
			return
		}

		// Tier 3: Per hour limit
		allowed, _, _ = storage.CheckAndIncrement(
			fmt.Sprintf("%s:hour:%d", clientIP, now.Unix()/3600),
			3600, // 1 hour window
			1000, // 1000 requests per hour
		)

		if !allowed {
			c.Header("X-RateLimit-Type", "per-hour")
			c.JSON(429, gin.H{"error": "Rate limit exceeded (per hour)"})
			c.Abort()
			return
		}

		c.Next()
	}
}
