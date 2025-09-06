package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nakku/api-gateway/config"
	"nakku/api-gateway/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup logger
	logger := setupLogger(cfg.Log.Level, cfg.Log.Format)

	// Create Gin router
	router := setupRouter(cfg, logger)

	// Create HTTP server
	server := &http.Server{
		Addr:         cfg.GetServerAddr(),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.WithField("addr", server.Addr).Info("Starting API Gateway server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shutdown")
	}

	logger.Info("Server exited")
}

// setupLogger configures the logger
func setupLogger(level, format string) *logrus.Logger {
	logger := logrus.New()

	// Set log level
	switch level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Set log format
	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return logger
}

// setupRouter configures the Gin router
func setupRouter(cfg *config.Config, logger *logrus.Logger) *gin.Engine {
	// Set Gin mode
	if cfg.Server.Host == "0.0.0.0" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Add middleware
	router.Use(middleware.Logging(middleware.LoggingConfig{
		Logger:   logger,
		SkipPath: []string{"/health", "/metrics"},
	}))
	router.Use(middleware.RequestID())
	router.Use(middleware.CORS(middleware.DefaultCORSConfig()))
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"service":   "api-gateway",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// User service routes
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/register", proxyToService(cfg, "user-service", "/users/register"))
			userRoutes.POST("/login", proxyToService(cfg, "user-service", "/users/login"))
			userRoutes.GET("/profile", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), proxyToService(cfg, "user-service", "/users/profile"))
			userRoutes.PUT("/profile", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), proxyToService(cfg, "user-service", "/users/profile"))
		}

		// Product service routes
		productRoutes := api.Group("/products")
		{
			productRoutes.GET("", proxyToService(cfg, "product-service", "/products"))
			productRoutes.GET("/:id", proxyToService(cfg, "product-service", "/products/:id"))
			productRoutes.GET("/search", proxyToService(cfg, "product-service", "/products/search"))
			productRoutes.GET("/category/:category", proxyToService(cfg, "product-service", "/products/category/:category"))
		}

		// Cart service routes
		cartRoutes := api.Group("/cart")
		cartRoutes.Use(middleware.Auth(middleware.AuthConfig{
			SecretKey: cfg.JWT.Secret,
			Logger:    logger,
		}))
		{
			cartRoutes.GET("", proxyToService(cfg, "cart-service", "/cart"))
			cartRoutes.POST("/items", proxyToService(cfg, "cart-service", "/cart/items"))
			cartRoutes.PUT("/items/:id", proxyToService(cfg, "cart-service", "/cart/items/:id"))
			cartRoutes.DELETE("/items/:id", proxyToService(cfg, "cart-service", "/cart/items/:id"))
			cartRoutes.DELETE("", proxyToService(cfg, "cart-service", "/cart"))
		}

		// Order service routes
		orderRoutes := api.Group("/orders")
		orderRoutes.Use(middleware.Auth(middleware.AuthConfig{
			SecretKey: cfg.JWT.Secret,
			Logger:    logger,
		}))
		{
			orderRoutes.POST("", proxyToService(cfg, "order-service", "/orders"))
			orderRoutes.GET("", proxyToService(cfg, "order-service", "/orders"))
			orderRoutes.GET("/:id", proxyToService(cfg, "order-service", "/orders/:id"))
			orderRoutes.PUT("/:id/cancel", proxyToService(cfg, "order-service", "/orders/:id/cancel"))
		}

		// Payment service routes
		paymentRoutes := api.Group("/payments")
		paymentRoutes.Use(middleware.Auth(middleware.AuthConfig{
			SecretKey: cfg.JWT.Secret,
			Logger:    logger,
		}))
		{
			paymentRoutes.POST("", proxyToService(cfg, "payment-service", "/payments"))
			paymentRoutes.GET("/:id", proxyToService(cfg, "payment-service", "/payments/:id"))
			paymentRoutes.POST("/:id/refund", proxyToService(cfg, "payment-service", "/payments/:id/refund"))
		}

		// Delivery service routes
		deliveryRoutes := api.Group("/delivery")
		deliveryRoutes.Use(middleware.Auth(middleware.AuthConfig{
			SecretKey: cfg.JWT.Secret,
			Logger:    logger,
		}))
		{
			deliveryRoutes.GET("/track/:id", proxyToService(cfg, "delivery-service", "/delivery/track/:id"))
			deliveryRoutes.GET("/history", proxyToService(cfg, "delivery-service", "/delivery/history"))
		}

		// Location service routes
		locationRoutes := api.Group("/location")
		{
			locationRoutes.POST("/validate", proxyToService(cfg, "location-service", "/location/validate"))
			locationRoutes.GET("/delivery-zones", proxyToService(cfg, "location-service", "/location/delivery-zones"))
		}

		// Analytics service routes (admin only)
		analyticsRoutes := api.Group("/analytics")
		analyticsRoutes.Use(middleware.Auth(middleware.AuthConfig{
			SecretKey: cfg.JWT.Secret,
			Logger:    logger,
		}))
		analyticsRoutes.Use(middleware.RequireRole("admin"))
		{
			analyticsRoutes.GET("/dashboard", proxyToService(cfg, "analytics-service", "/analytics/dashboard"))
			analyticsRoutes.GET("/reports", proxyToService(cfg, "analytics-service", "/analytics/reports"))
		}
	}

	return router
}

// proxyToService creates a proxy handler to forward requests to microservices
func proxyToService(cfg *config.Config, serviceName, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get service URL from environment or config
		serviceURL := getServiceURL(serviceName)
		if serviceURL == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   fmt.Sprintf("Service %s not available", serviceName),
			})
			return
		}

		// Forward request to the service
		forwardRequest(c, serviceURL, path)
	}
}

// getServiceURL returns the URL for a service
func getServiceURL(serviceName string) string {
	// In a real implementation, you might use service discovery
	// For now, we'll use environment variables or hardcoded URLs
	serviceURLs := map[string]string{
		"user-service":        getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		"product-service":     getEnv("PRODUCT_SERVICE_URL", "http://localhost:8082"),
		"inventory-service":   getEnv("INVENTORY_SERVICE_URL", "http://localhost:8083"),
		"order-service":       getEnv("ORDER_SERVICE_URL", "http://localhost:8084"),
		"cart-service":        getEnv("CART_SERVICE_URL", "http://localhost:8085"),
		"payment-service":     getEnv("PAYMENT_SERVICE_URL", "http://localhost:8086"),
		"delivery-service":    getEnv("DELIVERY_SERVICE_URL", "http://localhost:8087"),
		"notification-service": getEnv("NOTIFICATION_SERVICE_URL", "http://localhost:8088"),
		"location-service":    getEnv("LOCATION_SERVICE_URL", "http://localhost:8089"),
		"analytics-service":   getEnv("ANALYTICS_SERVICE_URL", "http://localhost:8090"),
	}

	return serviceURLs[serviceName]
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// forwardRequest forwards the request to the target service
func forwardRequest(c *gin.Context, targetURL, path string) {
	// This is a simplified proxy implementation
	// In a real implementation, you would use a proper HTTP client
	// to forward the request with all headers, body, etc.

	c.JSON(http.StatusNotImplemented, gin.H{
		"success": false,
		"error":   "Proxy functionality not implemented yet",
		"target":  targetURL + path,
	})
}
