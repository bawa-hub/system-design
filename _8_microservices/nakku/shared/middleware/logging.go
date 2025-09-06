package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggingConfig holds logging middleware configuration
type LoggingConfig struct {
	Logger   *logrus.Logger
	SkipPath []string
}

// Logging returns a logging middleware
func Logging(config LoggingConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip logging for specified paths
		for _, path := range config.SkipPath {
			if c.Request.URL.Path == path {
				c.Next()
				return
			}
		}

		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get client IP
		clientIP := c.ClientIP()

		// Get method and status
		method := c.Request.Method
		statusCode := c.Writer.Status()

		// Get body size
		bodySize := c.Writer.Size()

		// Create log entry
		entry := config.Logger.WithFields(logrus.Fields{
			"status":     statusCode,
			"latency":    latency,
			"client_ip":  clientIP,
			"method":     method,
			"path":       path,
			"body_size":  bodySize,
			"user_agent": c.Request.UserAgent(),
		})

		// Add query parameters if present
		if raw != "" {
			entry = entry.WithField("query", raw)
		}

		// Add request ID if present
		if requestID := c.GetHeader("X-Request-ID"); requestID != "" {
			entry = entry.WithField("request_id", requestID)
		}

		// Add user ID if present
		if userID := c.GetString("user_id"); userID != "" {
			entry = entry.WithField("user_id", userID)
		}

		// Log based on status code
		switch {
		case statusCode >= 500:
			entry.Error("Server error")
		case statusCode >= 400:
			entry.Warn("Client error")
		default:
			entry.Info("Request completed")
		}
	}
}

// RequestID adds a request ID to the context
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)
		c.Next()
	}
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString generates a random string of specified length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
