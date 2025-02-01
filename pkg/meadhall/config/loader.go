package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/adrg/xdg"
	yaml "gopkg.in/yaml.v3"
)

var (
	NoConfig error = errors.New("No configuration file exists")
)

// Loads the configuration from $XDG_CONFIG_PATH/meadhall/config.yaml.
// If it fails to load or parse, the method will return a valid object with sane defaults
// so it won't breka the application
func LoadConfig() (YamlConfig, error) {
	base := YamlConfig{}
	configPath, err := xdg.ConfigFile("meadhall/config.yaml")

	if err != nil {
		slog.Warn("Unable to find config file", "error", err)
		return base, errors.Join(NoConfig, err)
	}

	data, err := os.ReadFile(configPath)

	if err != nil {
		slog.Warn("Failed to read config", "error", err, "config", configPath)
		return base, err
	}

	err = yaml.Unmarshal(data, &base)

	if err != nil {
		slog.Warn("Failed to parse config", "error", err, "config", configPath)
		return base, err
	}

	return base, nil
}
