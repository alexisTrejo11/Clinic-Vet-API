// File: app/config/settings.go
package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// AppSettings holds all application configuration
type AppSettings struct {
	// Server Configuration
	Server ServerConfig `json:"server"`

	// Database Configuration
	Database DatabaseConfig `json:"database"`

	// Redis Configuration
	Redis RedisConfig `json:"redis"`

	// Authentication Configuration
	Auth AuthConfig `json:"auth"`

	// External Services Configuration
	Services ServicesConfig `json:"services"`

	// Rate Limiting Configuration
	RateLimit RateLimitConfig `json:"rate_limit"`

	// CORS Configuration
	CORS CORSConfig `json:"cors"`

	// Application Configuration
	App AppConfig `json:"app"`
}

type ServerConfig struct {
	Port         string        `json:"port"`
	Host         string        `json:"host"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	Environment  string        `json:"environment"`
}

type AuthConfig struct {
	JWTSecret          string        `json:"jwt_secret"`
	JWTExpirationTime  time.Duration `json:"jwt_expiration_time"`
	RefreshTokenExpiry time.Duration `json:"refresh_token_expiry"`
	PasswordMinLength  int           `json:"password_min_length"`
	MaxLoginAttempts   int           `json:"max_login_attempts"`
	LockoutDuration    time.Duration `json:"lockout_duration"`
}

type ServicesConfig struct {
	Twilio TwilioConfig `json:"twilio"`
	Email  EmailConfig  `json:"email"`
	Mongo  MongoConfig  `json:"mongo"`
}

type CORSConfig struct {
	AllowOrigins     []string      `json:"allow_origins"`
	AllowMethods     []string      `json:"allow_methods"`
	AllowHeaders     []string      `json:"allow_headers"`
	AllowCredentials bool          `json:"allow_credentials"`
	MaxAge           time.Duration `json:"max_age"`
}

type AppConfig struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	Debug         bool   `json:"debug"`
	LogLevel      string `json:"log_level"`
	EnableSwagger bool   `json:"enable_swagger"`
}

// LoadSettings loads and validates all application settings
func LoadSettings() (*AppSettings, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// Not fatal if .env doesn't exist
		fmt.Printf("Warning: Could not load .env file: %v\n", err)
	}

	settings := &AppSettings{}

	// Load and validate each configuration section
	if err := loadServerConfig(&settings.Server); err != nil {
		return nil, fmt.Errorf("server config error: %w", err)
	}

	if err := loadDatabaseConfig(&settings.Database); err != nil {
		return nil, fmt.Errorf("database config error: %w", err)
	}

	if err := loadRedisConfig(&settings.Redis); err != nil {
		return nil, fmt.Errorf("redis config error: %w", err)
	}

	if err := loadAuthConfig(&settings.Auth); err != nil {
		return nil, fmt.Errorf("auth config error: %w", err)
	}

	if err := loadServicesConfig(&settings.Services); err != nil {
		return nil, fmt.Errorf("services config error: %w", err)
	}

	loadRateLimitConfig(&settings.RateLimit)
	loadCORSConfig(&settings.CORS)
	loadAppConfig(&settings.App)

	return settings, nil
}

func loadServerConfig(config *ServerConfig) error {
	config.Port = getEnvWithDefault("SERVER_PORT", "8080")
	config.Host = getEnvWithDefault("SERVER_HOST", "0.0.0.0")
	config.Environment = getEnvWithDefault("ENVIRONMENT", "development")

	readTimeout, err := parseDuration("SERVER_READ_TIMEOUT", "30s")
	if err != nil {
		return fmt.Errorf("invalid SERVER_READ_TIMEOUT: %w", err)
	}
	config.ReadTimeout = readTimeout

	writeTimeout, err := parseDuration("SERVER_WRITE_TIMEOUT", "30s")
	if err != nil {
		return fmt.Errorf("invalid SERVER_WRITE_TIMEOUT: %w", err)
	}
	config.WriteTimeout = writeTimeout

	return nil
}

func loadDatabaseConfig(config *DatabaseConfig) error {
	config.URL = os.Getenv("DATABASE_URL")
	if config.URL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}

	var err error
	config.MaxOpenConns, err = parseIntWithDefault("DB_MAX_OPEN_CONNS", 25)
	if err != nil {
		return fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %w", err)
	}

	config.MaxIdleConns, err = parseIntWithDefault("DB_MAX_IDLE_CONNS", 5)
	if err != nil {
		return fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %w", err)
	}

	config.ConnMaxLifetime, err = parseDuration("DB_CONN_MAX_LIFETIME", "1h")
	if err != nil {
		return fmt.Errorf("invalid DB_CONN_MAX_LIFETIME: %w", err)
	}

	config.ConnMaxIdleTime, err = parseDuration("DB_CONN_MAX_IDLE_TIME", "30m")
	if err != nil {
		return fmt.Errorf("invalid DB_CONN_MAX_IDLE_TIME: %w", err)
	}

	return nil
}

func loadRedisConfig(config *RedisConfig) error {
	config.URL = os.Getenv("REDIS_URL")
	if config.URL == "" {
		return fmt.Errorf("REDIS_URL is required")
	}

	config.Password = getEnvWithDefault("REDIS_PASSWORD", "")

	var err error
	config.DB, err = parseIntWithDefault("REDIS_DB", 0)
	if err != nil {
		return fmt.Errorf("invalid REDIS_DB: %w", err)
	}

	return nil
}

func loadAuthConfig(config *AuthConfig) error {
	config.JWTSecret = os.Getenv("JWT_SECRET")
	if config.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	var err error
	config.JWTExpirationTime, err = parseDuration("JWT_EXPIRATION_TIME", "24h")
	if err != nil {
		return fmt.Errorf("invalid JWT_EXPIRATION_TIME: %w", err)
	}

	config.RefreshTokenExpiry, err = parseDuration("REFRESH_TOKEN_EXPIRY", "168h") // 7 days
	if err != nil {
		return fmt.Errorf("invalid REFRESH_TOKEN_EXPIRY: %w", err)
	}

	config.PasswordMinLength, err = parseIntWithDefault("PASSWORD_MIN_LENGTH", 8)
	if err != nil {
		return fmt.Errorf("invalid PASSWORD_MIN_LENGTH: %w", err)
	}

	config.MaxLoginAttempts, err = parseIntWithDefault("MAX_LOGIN_ATTEMPTS", 5)
	if err != nil {
		return fmt.Errorf("invalid MAX_LOGIN_ATTEMPTS: %w", err)
	}

	config.LockoutDuration, err = parseDuration("LOCKOUT_DURATION", "30m")
	if err != nil {
		return fmt.Errorf("invalid LOCKOUT_DURATION: %w", err)
	}

	return nil
}

func loadServicesConfig(config *ServicesConfig) error {
	// Twilio
	config.Twilio.AccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	config.Twilio.AuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	config.Twilio.FromPhone = os.Getenv("TWILIO_FROM_PHONE")

	if config.Twilio.AccountSID == "" || config.Twilio.AuthToken == "" {
		return fmt.Errorf("TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN are required")
	}

	// Email
	config.Email.SMTPHost = os.Getenv("SMTP_HOST")
	config.Email.SMTPUsername = os.Getenv("SMTP_USERNAME")
	config.Email.SMTPPassword = os.Getenv("SMTP_PASSWORD")
	config.Email.FromEmail = os.Getenv("FROM_EMAIL")
	config.Email.FromName = getEnvWithDefault("FROM_NAME", "Clinic Vet API")

	if config.Email.SMTPHost == "" {
		return fmt.Errorf("SMTP_HOST is required")
	}

	var err error
	config.Email.SMTPPort, err = parseIntWithDefault("SMTP_PORT", 587)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %w", err)
	}

	// MongoDB
	config.Mongo.URI = os.Getenv("MONGO_URI")
	config.Mongo.Database = getEnvWithDefault("MONGO_DATABASE", "clinic_vet")

	if config.Mongo.URI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}

	config.Mongo.Timeout, err = parseDuration("MONGO_TIMEOUT", "10s")
	if err != nil {
		return fmt.Errorf("invalid MONGO_TIMEOUT: %w", err)
	}

	return nil
}

func loadRateLimitConfig(config *RateLimitConfig) {
	config.Enabled = parseBoolWithDefault("RATE_LIMIT_ENABLED", true)
	config.RequestsPerSecond, _ = parseIntWithDefault("RATE_LIMIT_RPS", 100)
	config.WindowDuration, _ = parseDuration("RATE_LIMIT_WINDOW", "1m")
	config.BurstSize, _ = parseIntWithDefault("RATE_LIMIT_BURST", 20)
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

func loadAppConfig(config *AppConfig) {
	config.Name = getEnvWithDefault("APP_NAME", "Clinic Vet API")
	config.Version = getEnvWithDefault("APP_VERSION", "1.0.0")
	config.Debug = parseBoolWithDefault("DEBUG", false)
	config.LogLevel = getEnvWithDefault("LOG_LEVEL", "info")
	config.EnableSwagger = parseBoolWithDefault("ENABLE_SWAGGER", true)
}

// Validation methods
func (s *AppSettings) Validate() error {
	var errors []string

	// Validate server config
	if s.Server.Port == "" {
		errors = append(errors, "server port cannot be empty")
	}

	// Validate database config
	if s.Database.URL == "" {
		errors = append(errors, "database URL is required")
	}

	// Validate auth config
	if s.Auth.JWTSecret == "" {
		errors = append(errors, "JWT secret is required")
	}
	if len(s.Auth.JWTSecret) < 32 {
		errors = append(errors, "JWT secret should be at least 32 characters long")
	}

	// Validate services
	if s.Services.Twilio.AccountSID == "" {
		errors = append(errors, "Twilio Account SID is required")
	}
	if s.Services.Email.SMTPHost == "" {
		errors = append(errors, "SMTP host is required")
	}
	if s.Services.Mongo.URI == "" {
		errors = append(errors, "MongoDB URI is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("configuration validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// IsDevelopment returns true if running in development mode
func (s *AppSettings) IsDevelopment() bool {
	return s.Server.Environment == "development" || s.Server.Environment == "dev"
}

// IsProduction returns true if running in production mode
func (s *AppSettings) IsProduction() bool {
	return s.Server.Environment == "production" || s.Server.Environment == "prod"
}

// GetServerAddr returns the full server address
func (s *AppSettings) GetServerAddr() string {
	return s.Server.Host + ":" + s.Server.Port
}
