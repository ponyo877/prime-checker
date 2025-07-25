package config

import (
	"os"

	"github.com/ponyo877/prime-checker/internal/shared/infrastructure"
)

func LoadDatabaseConfig() infrastructure.DatabaseConfig {
	return infrastructure.DatabaseConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
}

func LoadMessagingConfig() infrastructure.MessagingConfig {
	return infrastructure.MessagingConfig{
		Host: os.Getenv("NATS_HOST"),
		Port: os.Getenv("NATS_PORT"),
	}
}
