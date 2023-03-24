package configs

import (
	"fmt"

	util "github.com/pratik1998/secret-share-backend/utils"
)

// DatabaseConfig is the configuration for the database.
type DatabaseConfig struct {
	Host     string        `json:"host"`
	Port     int           `json:"port"`
	User     string        `json:"user"`
	Password string        `json:"password"`
	Database string        `json:"database"`
	GetURL   func() string `json:"-"` // This is a function that returns the database URL
}

func getMongoDBURL() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", config.Database.User, config.Database.Password, config.Database.Host)
}

// Config is the configuration for the application.
type Config struct {
	Database DatabaseConfig
}

var config *Config

// NewConfig creates a new configuration for the application.
func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     util.GodotEnv("DB_HOST"),
			Port:     util.GodotEnvInt("DB_PORT"),
			User:     util.GodotEnv("DB_USER"),
			Password: util.GodotEnv("DB_PASSWORD"),
			Database: util.GodotEnv("DB_DATABASE"),
			GetURL:   getMongoDBURL,
		},
	}
}

// GetConfig returns the configuration for the application.
func GetConfig() *Config {
	if config == nil {
		config = NewConfig()
	}
	return config
}
