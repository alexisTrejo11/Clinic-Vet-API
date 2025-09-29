package ginutils

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginMetadata struct {
	IP         string
	UserAgent  string
	DeviceInfo string
	Timestamp  int64
}

func NewLoginMetadata(c *gin.Context) *LoginMetadata {
	ip := extractIP(c)
	userAgent := c.Request.UserAgent()

	return &LoginMetadata{
		IP:         ip,
		UserAgent:  userAgent,
		DeviceInfo: extractDeviceInfo(userAgent),
		Timestamp:  time.Now().Unix(),
	}
}

func extractIP(c *gin.Context) string {
	if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := c.GetHeader("X-Real-IP"); ip != "" {
		return ip
	}

	return c.ClientIP()
}

func extractDeviceInfo(userAgent string) string {
	ua := strings.ToLower(userAgent)

	switch {
	case strings.Contains(ua, "mobile"):
		return "mobile"
	case strings.Contains(ua, "tablet"):
		return "tablet"
	default:
		return "desktop"
	}
}

func (lm *LoginMetadata) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ip":         lm.IP,
		"user_agent": lm.UserAgent,
		"device":     lm.DeviceInfo,
		"timestamp":  lm.Timestamp,
	}
}
