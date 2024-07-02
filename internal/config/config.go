package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
)

func New() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("env.Parse: %w", err)
	}

	return cfg, nil
}

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS,notEmpty"`
	IllustratorConfig
}

type IllustratorConfig struct {
	GroupMaxSentences int `env:"GROUP_MAX_SENTENCES,notEmpty"`
}
