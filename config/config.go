package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
