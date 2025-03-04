package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"time"
)

var AppConfig *Config

type Config struct {
	ServerConfig ServerConfig
	DBConfig     DBConfig
}

type DBConfig struct {
	DBDriver     string        `env:"DB_DRIVER,required"`
	DBSource     string        `env:"DB_SOURCE,required"`
	DbHost       string        `env:"DB_HOST,required"`
	DbPort       string        `env:"DB_PORT,required"`
	DbUser       string        `env:"DB_USER,required"`
	DbPassword   string        `env:"DB_PASSWORD,required"`
	DbName       string        `env:"DB_NAME,required"`
	DbSslMode    string        `env:"DB_SSLMODE,required"`
	MaxOpenConns int           `env:"DB_MAX_OPEN_CONNECTIONS,required"`
	MaxIdleConns int           `env:"DB_MAX_IDLE_CONNECTIONS,required"`
	MaxIdleTime  time.Duration `env:"DB_MAX_IDLE_TIME,required"`
	Timeout      time.Duration `env:"DB_TIMEOUT,required"`
}

type ServerConfig struct {
	Port    string `env:"SERVER_PORT,required"`
	Version string `env:"SERVER_VERSION,required"`
}

func LoadConfig() error {
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatal("Error loading app.env file")
	}

	config := &Config{}
	if err := env.Parse(config); err != nil {
		log.Fatal("Error parsing config")
	}

	serverConfig := &ServerConfig{}
	if err := env.Parse(serverConfig); err != nil {
		log.Fatal("Error parsing config")
	}

	AppConfig = config

	return nil
}
