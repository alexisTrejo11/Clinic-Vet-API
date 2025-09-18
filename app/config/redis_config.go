package config

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	URL      string `json:"url"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var RedisClient *redis.Client

func InitRedis(redisUrl string) {
	if redisUrl == "" {
		panic("REDIS_URL environment variable is not set")
	}

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(opt)
}
