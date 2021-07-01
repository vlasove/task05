package store

import (
	"fmt"

	"github.com/vlasove/test05/internal/app/helper"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Host:     helper.EnvLoader("host", "127.0.0.1"),
		Port:     helper.EnvLoader("port", "5432"),
		User:     helper.EnvLoader("user", "postgres"),
		Password: helper.EnvLoader("password", "postgres"),
		DBName:   helper.EnvLoader("dbname", "postgresdb"),
		SSLMode:  helper.EnvLoader("sslmode", "disable"),
	}
}

// buildConnStr ...
func (c *Config) buildConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}
