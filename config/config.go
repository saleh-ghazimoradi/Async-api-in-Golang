package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Env string

const (
	EnvTest Env = "test"
	EnvDev  Env = "dev"
)

type Config struct {
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabasePortTest string `env:"DATABASE_PORT_TEST"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	ENV              Env    `env:"ENV" envDefault:"dev"`
	ProjectRoot      string `env:"PROJECT_ROOT"`
}

func (c *Config) DatabaseURL() string {
	port := c.DatabasePort
	if c.ENV == EnvTest {
		port = c.DatabasePortTest
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, port, c.DatabaseName)
}

func NewConfig() (*Config, error) {
	var cfg Config
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
