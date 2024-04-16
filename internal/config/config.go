package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	Application Application `mapstructure:"application"`
	Database    Database    `mapstructure:"database"`
	Cache       Cache       `mapstructure:"cache"`
	Logger      Logger      `mapstructure:"logger"`
}

// New ...
func New() *Config {
	cfg, err := load()
	if err != nil {
		panic(err)
	}
	return cfg
}

// load ...
// loadConfigFile attempts to load the configuration file from the given path.
func loadConfigFile(path string) error {
	viper.Reset()

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}

// findConfigPath finds the configuration file by traversing directories upwards.
func findConfigPath() (string, error) {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}

	for i := 0; i < 10; i++ {
		configPath := filepath.Join(currentDir, "resource")
		if err := loadConfigFile(configPath); err == nil {
			return configPath, nil
		}
		currentDir = filepath.Dir(currentDir)
	}

	return "", fmt.Errorf("config file not found")
}

// load reads the configuration from the config file using Viper.
func load() (*Config, error) {
	configPath, err := findConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to find config file: %v", err)
	}

	// Set default values
	viper.SetDefault("application.name", "MyApp")
	viper.SetDefault("application.version", "1.0.0")
	viper.SetDefault("application.address", "localhost")
	viper.SetDefault("application.port", 3000)

	if err := loadConfigFile(configPath); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return cfg, nil
}
