package config

import (
	"os"
	"testing"
)

const dummyOpenAIKey = "dummyOpenAIKey"

func TestLoadConfig_Success(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", dummyOpenAIKey)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.OPENAI_API_KEY != dummyOpenAIKey {
		t.Errorf("Expected OPENAI_API_KEY to be %s, but got %s", dummyOpenAIKey, cfg.OPENAI_API_KEY)
	}
}

func TestLoadConfig_MissingEnv(t *testing.T) {
	os.Unsetenv("OPENAI_API_KEY")

	cfg, err := LoadConfig()
	if err == nil {
		t.Fatalf("Error wasn't returned. cfg: %+v", cfg)
	}

	expectedErrString := "OPENAI_API_KEY env variable is required."
	if err.Error() != expectedErrString {
		t.Errorf("Expected error message to be %s, but got %s", expectedErrString, err.Error())
	}
}
