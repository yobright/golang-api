package config

import "fmt"
import "github.com/caarlos0/env/v11"

type Config struct {
	DatabaseName     string `env:"DB_NAME"`
	DatabaseUser     string `env:"DB_USER"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     string `env:"DB_PORT"`
	DatabasePassword string `env:"DB_PASSWORD"`
}

func New() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("error parsing env variables: %v", err)
	}

	return &cfg, nil
}
