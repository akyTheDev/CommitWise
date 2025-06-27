package config

import (
	"errors"
	"os"
)

type Config struct {
	OPENAI_API_KEY string
}

func LoadConfig() (*Config, error) {
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	if openAiApiKey == "" {
		return nil, errors.New("OPENAI_API_KEY env variable is required.")
	}

	return &Config{
		OPENAI_API_KEY: openAiApiKey,
	}, nil
}
