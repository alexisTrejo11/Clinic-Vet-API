package config

import "time"

type RateLimitConfig struct {
	Enabled           bool          `json:"enabled"`
	RequestsPerSecond int           `json:"requests_per_second"`
	WindowDuration    time.Duration `json:"window_duration"`
	BurstSize         int           `json:"burst_size"`
}
