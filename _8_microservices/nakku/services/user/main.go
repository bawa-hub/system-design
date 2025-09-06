package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nakku/services/user/config"
	"nakku/services/user/database"
	"nakku/services/user/middleware"
	"nakku/services/user/utils"
	"nakku/services/user/handlers"
	userModels "nakku/services/user/models"
	"nakku/services/user/repository"
	"nakku/services/user/service"

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

	// Connect to database
	if err := database.Connect(&cfg.Database); err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
	}
	defer database.Close()

	// Run database migrations
	if err := database.Migrate(&userModels.User{}, &userModels.UserAddress{}); err != nil {
		logger.WithError(err).Fatal("Failed to run database migrations")
	}

	// Setup Kafka producer
	producer := utils.NewKafkaProducer(cfg.Kafka.Brokers, logger)
	defer producer.Close()

	// Initialize repository
	userRepo := repository.NewUserRepository(database.GetDB())

	// Initialize service
	userService := service.NewUserService(userRepo, producer, logger)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService, logger)

	// Setup router
	router := setupRouter(cfg, logger, userHandler)

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
		logger.WithField("addr", server.Addr).Info("Starting User Service server")
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
func setupRouter(cfg *config.Config, logger *logrus.Logger, userHandler *handlers.UserHandler) *gin.Engine {
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
			"service":   "user-service",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
			users.GET("/profile", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), userHandler.GetProfile)
			users.PUT("/profile", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), userHandler.UpdateProfile)
			users.POST("/addresses", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), userHandler.AddAddress)
			users.GET("/addresses", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), userHandler.GetAddresses)
			users.PUT("/addresses/:id", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), userHandler.UpdateAddress)
			users.DELETE("/addresses/:id", middleware.Auth(middleware.AuthConfig{
				SecretKey: cfg.JWT.Secret,
				Logger:    logger,
			}), userHandler.DeleteAddress)
		}
	}

	return router
}
