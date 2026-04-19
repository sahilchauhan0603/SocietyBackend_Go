package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
	SMTP     SMTPConfig
	Admin    AdminConfig
}

type ServerConfig struct {
	Port            string
	AllowedOrigins  []string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

type AuthConfig struct {
	JWTKey       string
	ClientID     string
	ClientSecret string
	TenantID     string
	RedirectURL  string
}

type SMTPConfig struct {
	Server   string
	Port     string
	User     string
	Password string
}

type AdminConfig struct {
	User     string
	Password string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	required := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASS",
		"JWT_KEY",
		"ADMIN_USER",
		"ADMIN_PASS",
		"SMTP_SERVER",
		"SMTP_PORT",
		"EMAIL_USER",
		"EMAIL_PASSWORD",
		"CLIENT_ID",
		"CLIENT_SECRET",
		"TENANT_ID",
		"REDIRECT_URL",
	}

	missing := make([]string, 0)
	for _, key := range required {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	cfg := &Config{
		Server: ServerConfig{
			Port:            getEnvWithDefault("PORT", "8080"),
			AllowedOrigins:  parseCSVEnvWithDefault("CORS_ALLOWED_ORIGINS", []string{"https://societymanagementfrontend-h3v3.onrender.com", "http://localhost:8000", "http://localhost:5173"}),
			ReadTimeout:     parseDurationWithDefault("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout:    parseDurationWithDefault("SERVER_WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:     parseDurationWithDefault("SERVER_IDLE_TIMEOUT", 60*time.Second),
			ShutdownTimeout: parseDurationWithDefault("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			SSLMode:  getEnvWithDefault("DB_SSLMODE", "verify-full"),
		},
		Auth: AuthConfig{
			JWTKey:       os.Getenv("JWT_KEY"),
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			TenantID:     os.Getenv("TENANT_ID"),
			RedirectURL:  os.Getenv("REDIRECT_URL"),
		},
		SMTP: SMTPConfig{
			Server:   os.Getenv("SMTP_SERVER"),
			Port:     os.Getenv("SMTP_PORT"),
			User:     os.Getenv("EMAIL_USER"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		},
		Admin: AdminConfig{
			User:     os.Getenv("ADMIN_USER"),
			Password: os.Getenv("ADMIN_PASS"),
		},
	}

	return cfg, nil
}

func parseDurationWithDefault(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func parseCSVEnvWithDefault(key string, fallback []string) []string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parts := strings.Split(value, ",")
	origins := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			origins = append(origins, trimmed)
		}
	}

	if len(origins) == 0 {
		return fallback
	}

	return origins
}

func getEnvWithDefault(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}
