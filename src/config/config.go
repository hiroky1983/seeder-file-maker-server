package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env          string `envconfig:"ENVIRONMENT" required:"true"`
	AppName      string `envconfig:"APP_NAME" required:"true"`
	AppPort      string `envconfig:"APP_PORT" required:"true" default:"8080"`
	OpenAiAPIKey string `envconfig:"OPEN_AI_API_KEY" required:"true"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
