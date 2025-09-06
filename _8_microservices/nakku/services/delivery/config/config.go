package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Log      LogConfig      `mapstructure:"log"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	IdleTimeout  int    `mapstructure:"idle_timeout"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
	MaxConns int    `mapstructure:"max_conns"`
	MinConns int    `mapstructure:"min_conns"`
}

// KafkaConfig holds Kafka configuration
type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
	GroupID string   `mapstructure:"group_id"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// LoadConfig loads configuration from environment variables and config files
func LoadConfig() (*Config, error) {
	config := &Config{}

	// Set default values
	setDefaults()

	// Load from environment variables
	loadFromEnv(config)

	// Load from config file if exists
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./shared/config")

	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(config); err != nil {
			return nil, fmt.Errorf("failed to unmarshal config: %w", err)
		}
	}

	return config, nil
}

// setDefaults sets default configuration values
func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.read_timeout", 30)
	viper.SetDefault("server.write_timeout", 30)
	viper.SetDefault("server.idle_timeout", 120)

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.name", "nakku")
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.max_conns", 25)
	viper.SetDefault("database.min_conns", 5)

	viper.SetDefault("kafka.brokers", []string{"localhost:9092"})
	viper.SetDefault("kafka.group_id", "nakku-group")

	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	viper.SetDefault("jwt.secret", "nakku-secret-key")
	viper.SetDefault("jwt.expiration", 3600)

	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(config *Config) {
	// Server config
	if port := os.Getenv("PORT"); port != "" {
		viper.Set("server.port", port)
	}
	if host := os.Getenv("HOST"); host != "" {
		viper.Set("server.host", host)
	}

	// Database config
	if host := os.Getenv("DB_HOST"); host != "" {
		viper.Set("database.host", host)
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		viper.Set("database.port", port)
	}
	if user := os.Getenv("DB_USER"); user != "" {
		viper.Set("database.user", user)
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		viper.Set("database.password", password)
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		viper.Set("database.name", name)
	}

	// Kafka config
	if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		viper.Set("kafka.brokers", strings.Split(brokers, ","))
	}
	if groupID := os.Getenv("KAFKA_GROUP_ID"); groupID != "" {
		viper.Set("kafka.group_id", groupID)
	}

	// Redis config
	if url := os.Getenv("REDIS_URL"); url != "" {
		// Parse Redis URL (format: redis://host:port/db)
		if strings.HasPrefix(url, "redis://") {
			url = strings.TrimPrefix(url, "redis://")
			parts := strings.Split(url, "/")
			if len(parts) >= 2 {
				hostPort := parts[0]
				dbStr := parts[1]
				if db, err := strconv.Atoi(dbStr); err == nil {
					viper.Set("redis.db", db)
				}
				if hostPort != "" {
					hostPortParts := strings.Split(hostPort, ":")
					if len(hostPortParts) >= 1 {
						viper.Set("redis.host", hostPortParts[0])
					}
					if len(hostPortParts) >= 2 {
						viper.Set("redis.port", hostPortParts[1])
					}
				}
			}
		}
	}

	// JWT config
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		viper.Set("jwt.secret", secret)
	}

	// Log config
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		viper.Set("log.level", level)
	}
}

// GetDSN returns the database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}

// GetRedisAddr returns the Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
}

// GetServerAddr returns the server address
func (c *Config) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}
