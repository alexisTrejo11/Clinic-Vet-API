package log

import "go.uber.org/zap"

// Debug logs debug message with optional fields
func Debug(msg string, fields ...zap.Field) {
	App.Debug(msg, fields...)
}

// Info logs info message with optional fields
func Info(msg string, fields ...zap.Field) {
	App.Info(msg, fields...)
}

// Warn logs warning message with optional fields
func Warn(msg string, fields ...zap.Field) {
	App.Warn(msg, fields...)
}

// Error logs error message with optional fields
func Error(msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	App.Error(msg, fields...)
}

// WithCorrelation creates correlation_id field
func WithCorrelation(id string) zap.Field {
	return zap.String(CorrelationIDKey, id)
}

// WithUser creates user_id field
func WithUser(userID string) zap.Field {
	return zap.String(UserIDKey, userID)
}

// WithOperation creates operation field
func WithOperation(operation string) zap.Field {
	return zap.String(OperationKey, operation)
}

// WithEntity creates entity and entity_id fields
func WithEntity(entity, entityID string) []zap.Field {
	return []zap.Field{
		zap.String(EntityKey, entity),
		zap.String(EntityIDKey, entityID),
	}
}

// WithRequest creates HTTP request fields
func WithRequest(method, path string) []zap.Field {
	return []zap.Field{
		zap.String("http_method", method),
		zap.String("http_path", path),
	}
}
