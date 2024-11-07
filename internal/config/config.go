package config

import (
	"encoding/json"
	"fmt"
	"log"
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

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handler map[string]func(*State, Command) error
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

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		log.Fatalf("usage: login <username>")
	}
	s.Config.CurrentUserName = cmd.Args[0]

	fmt.Println("You are now logged in as", cmd.Args[0])

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.Handler == nil {
		c.Handler = make(map[string]func(*State, Command) error)
	}
	c.Handler[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handler[cmd.Name]
	if !ok {
		log.Fatalf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
