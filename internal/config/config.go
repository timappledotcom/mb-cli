package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	MicroBlogToken string `json:"micro_blog_token"`
}

func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	
	configPath := filepath.Join(home, ".config", "mb", "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Return empty config if not found, to prompt user later possibly
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) Save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	
	dir := filepath.Join(home, ".config", "mb")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file := filepath.Join(dir, "config.json")
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0600)
}
