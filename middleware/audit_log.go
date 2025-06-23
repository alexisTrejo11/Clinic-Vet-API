package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func AuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		var requestBody bytes.Buffer
		if c.Request.Body != nil {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				config.AppLogger.Error("Error reading request body", zap.Error(err))
				c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
				return
			}

			requestBody.Write(body)

			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		c.Next()

		duration := time.Since(start)
		config.AuditLogger.Info("audit_log",
			zap.String("request_id", requestID),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("request_body", requestBody.String()),
			zap.String("response_body", blw.body.String()),
			zap.Any("headers", c.Request.Header),
			zap.Any("errors", c.Errors),
		)

		config.AppLogger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
		)
	}
}
