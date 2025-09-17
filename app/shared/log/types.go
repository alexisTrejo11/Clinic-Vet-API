// shared/log/types.go
package log

import "go.uber.org/zap"

// Context keys
const (
	CorrelationIDKey = "correlation_id"
	UserIDKey        = "user_id"
	OperationKey     = "operation"
	EntityKey        = "entity"
	EntityIDKey      = "entity_id"
)

var (
	App   *zap.Logger
	Audit *zap.Logger
)

type LogLevel string

const (
	LevelDebug LogLevel = "DEBUG"
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelError LogLevel = "ERROR"
)
