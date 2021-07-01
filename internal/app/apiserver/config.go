package apiserver

import (
	"github.com/vlasove/test05/internal/app/helper"
	"github.com/vlasove/test05/internal/app/store"
)

// Config ...
type Config struct {
	BindAddr string
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: helper.EnvLoader("bind_addr", ":8000"),
		Store:    store.NewConfig(),
	}
}
