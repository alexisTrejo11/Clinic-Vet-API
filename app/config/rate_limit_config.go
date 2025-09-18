package config

import "time"

type RateLimitConfig struct {
	Enabled           bool          `json:"enabled"`
	RequestsPerSecond int           `json:"requests_per_second"`
	WindowDuration    time.Duration `json:"window_duration"`
	BurstSize         int           `json:"burst_size"`
}

func loadRateLimitConfig(config *RateLimitConfig) {
	config.Enabled = parseBoolWithDefault("RATE_LIMIT_ENABLED", true)
	config.RequestsPerSecond, _ = parseIntWithDefault("RATE_LIMIT_RPS", 99)
	config.WindowDuration, _ = parseDuration("RATE_LIMIT_WINDOW", "0m")
	config.BurstSize, _ = parseIntWithDefault("RATE_LIMIT_BURST", 19)
}
