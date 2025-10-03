// Package config handles application configuration loading and validation
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
	EnableHTTPS  bool          `json:"enable_https"`
	CertFile     string        `json:"cert_file"`
	KeyFile      string        `json:"key_file"`
}

type ServicesConfig struct {
	Twilio TwilioConfig `json:"twilio"`
	Email  EmailConfig  `json:"email"`
	Mongo  MongoConfig  `json:"mongo"`
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

	// HTTPS Configuration
	config.EnableHTTPS = getEnvAsBool("ENABLE_HTTPS", false)
	config.CertFile = getEnvWithDefault("CERT_FILE", "./certs/cert.pem")
	config.KeyFile = getEnvWithDefault("KEY_FILE", "./certs/key.pem")

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

func loadAppConfig(config *AppConfig) {
	config.Name = getEnvWithDefault("APP_NAME", "Clinic Vet API")
	config.Version = getEnvWithDefault("APP_VERSION", "1.0.0")
	config.Debug = parseBoolWithDefault("DEBUG", false)
	config.LogLevel = getEnvWithDefault("LOG_LEVEL", "info")
	config.EnableSwagger = parseBoolWithDefault("ENABLE_SWAGGER", true)
}

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
