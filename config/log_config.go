package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	AppLogger   *zap.Logger
	AuditLogger *zap.Logger
)

func InitLogger() {
	// App
	appConfig := zap.NewDevelopmentEncoderConfig()
	appConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	appConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	appConfig.TimeKey = "time"
	appConfig.LevelKey = "level"
	appConfig.MessageKey = "msg"

	consoleEncoder := zapcore.NewConsoleEncoder(appConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	AppLogger = zap.New(core, zap.AddCaller())

	// AUDIT
	auditFile := &lumberjack.Logger{
		Filename:   "logs/audit.log",
		MaxSize:    100, // MB
		MaxBackups: 30,
		MaxAge:     90, // days
		Compress:   true,
	}

	auditConfig := zap.NewProductionEncoderConfig()
	auditConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	auditConfig.TimeKey = "timestamp"
	auditConfig.MessageKey = "message"

	fileEncoder := zapcore.NewJSONEncoder(auditConfig)

	auditCore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(auditFile), zapcore.InfoLevel),
	)

	AuditLogger = zap.New(auditCore)
}

func SyncLogger() {
	_ = AppLogger.Sync()
	_ = AuditLogger.Sync()
}
