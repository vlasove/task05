package apiserver

import (
	"github.com/vlasove/test05/internal/app/helper"
)

// Config ...
type Config struct {
	BindAddr          string
	DatabaseConnector *DBConnector
	DatabaseURL       string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:          helper.EnvLoader("bind_addr", ":8000"),
		DatabaseConnector: NewDBConnector(),
	}
}
