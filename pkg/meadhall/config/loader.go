package config

import (
	"os"

	"github.com/adrg/xdg"
	yaml "gopkg.in/yaml.v3"
)

func LoadConfig() YamlConfig {

	configPath, err := xdg.ConfigFile("meadhall/config.yaml")

	if err != nil {
		panic("No configuration to read!")
	}

	data, err := os.ReadFile(configPath)

	if err != nil {
		panic("Failed to read config")
	}

	base := YamlConfig{}

	err = yaml.Unmarshal(data, &base)

	if err != nil {
		panic(err)
	}

	return base
}
