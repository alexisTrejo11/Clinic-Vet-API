package config

import (
	"strings"
	"time"
)

type CORSConfig struct {
	AllowOrigins     []string      `json:"allow_origins"`
	AllowMethods     []string      `json:"allow_methods"`
	AllowHeaders     []string      `json:"allow_headers"`
	AllowCredentials bool          `json:"allow_credentials"`
	MaxAge           time.Duration `json:"max_age"`
}

func loadCORSConfig(config *CORSConfig) {
	originsStr := getEnvWithDefault("CORS_ALLOW_ORIGINS", "*")
	if originsStr == "*" {
		config.AllowOrigins = []string{"*"}
	} else {
		config.AllowOrigins = strings.Split(originsStr, ",")
		// Trim whitespace
		for i, origin := range config.AllowOrigins {
			config.AllowOrigins[i] = strings.TrimSpace(origin)
		}
	}

	methodsStr := getEnvWithDefault("CORS_ALLOW_METHODS", "GET,POST,PUT,DELETE,OPTIONS")
	config.AllowMethods = strings.Split(methodsStr, ",")
	for i, method := range config.AllowMethods {
		config.AllowMethods[i] = strings.TrimSpace(method)
	}

	headersStr := getEnvWithDefault("CORS_ALLOW_HEADERS", "Origin,Content-Type,Authorization")
	config.AllowHeaders = strings.Split(headersStr, ",")
	for i, header := range config.AllowHeaders {
		config.AllowHeaders[i] = strings.TrimSpace(header)
	}

	config.AllowCredentials = parseBoolWithDefault("CORS_ALLOW_CREDENTIALS", true)
	config.MaxAge, _ = parseDuration("CORS_MAX_AGE", "12h")
}
