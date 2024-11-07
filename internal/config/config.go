package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
type State struct {
	Config *Config
}

const configFileName = ".gatorconfig.json"

// LoadConfig loads the configuration from a JSON file in the user's home directory.
func LoadConfig() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("error getting user home directory: %w", err)
	}
	filePath := filepath.Join(homeDir, configFileName)
	var cfg Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config JSON: %w", err)
	}
	return cfg, nil
}

func SetUser(cfg *Config, username string) error {
	cfg.CurrentUserName = username

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshalling config JSON: %w", err)
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}
	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
