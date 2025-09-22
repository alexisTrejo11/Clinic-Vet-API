package event

import (
	"clinic-vet-api/app/shared/log"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type EventLogger struct {
	logger *zap.Logger
}

func NewEventLogger() *EventLogger {
	return &EventLogger{
		logger: log.App,
	}
}

// LogEvent initializes logging for a new event and returns a context with event metadata
type ctxKey string

func (el *EventLogger) LogEvent(ctx context.Context, eventName, operation string, fields ...zap.Field) context.Context {
	eventID := generateEventID()
	ctx = context.WithValue(ctx, ctxKey("event_id"), eventID)
	ctx = context.WithValue(ctx, ctxKey("event_name"), eventName)
	ctx = context.WithValue(ctx, ctxKey("operation"), operation)

	allFields := append([]zap.Field{
		zap.String("event_id", eventID),
		zap.String("event_name", eventName),
		zap.String("operation", operation),
		zap.String("stage", "started"),
	}, fields...)

	el.logger.Info(fmt.Sprintf("Event: %s started", eventName), allFields...)
	return ctx
}

// LogOperationSuccess logs a successful operation
func (el *EventLogger) LogOperationSuccess(ctx context.Context, message string, additionalFields ...zap.Field) {
	el.logOperation(ctx, "success", message, additionalFields...)
}

// LogOperationError logs an error in an operation
func (el *EventLogger) LogOperationError(ctx context.Context, err error, message string, additionalFields ...zap.Field) {
	fields := append(additionalFields, zap.String("stage", "failed"), zap.Error(err))
	el.logOperation(ctx, "error", message, fields...)
}

// LogOperationWarning logs a warning in an operation
func (el *EventLogger) LogOperationWarning(ctx context.Context, message string, additionalFields ...zap.Field) {
	el.logOperation(ctx, "warning", message, additionalFields...)
}

// logOperation is a helper to log operations with a specific level
func (el *EventLogger) logOperation(ctx context.Context, level, message string, additionalFields ...zap.Field) {
	baseFields := el.getBaseFields(ctx)
	allFields := append(baseFields, additionalFields...)

	switch level {
	case "success":
		el.logger.Info(fmt.Sprintf("Success: %s", message), allFields...)
	case "error":
		el.logger.Error(fmt.Sprintf("Error: %s", message), allFields...)
	case "warning":
		el.logger.Warn(fmt.Sprintf("Warning: %s", message), allFields...)
	}
}

// getBaseFields extracts common fields from the context for logging
func (el *EventLogger) getBaseFields(ctx context.Context) []zap.Field {
	fields := []zap.Field{
		zap.String("event_id", getContextString(ctx, "event_id")),
		zap.String("event_name", getContextString(ctx, "event_name")),
		zap.String("operation", getContextString(ctx, "operation")),
	}

	if userID := getContextString(ctx, "user_id"); userID != "unknown" {
		fields = append(fields, zap.String("user_id", userID))
	}

	return fields
}

func generateEventID() string {
	return fmt.Sprintf("evt_%s_%s", time.Now().Format("20060102_150405"), randomString(6))
}

func getContextString(ctx context.Context, key string) string {
	if ctx == nil {
		return "unknown"
	}
	if value, ok := ctx.Value(ctxKey(key)).(string); ok {
		return value
	}
	return "unknown"
}

func randomString(length int) string {
	return "abc123"
}
