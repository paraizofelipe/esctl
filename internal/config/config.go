package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

const CONFIG_PATH = "%s/.config/esctl/config.toml"

type ConfigFile struct {
	Elastic  []string `toml:"elastic"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
}

func Load() (*ConfigFile, error) {
	path := fmt.Sprintf(CONFIG_PATH, os.Getenv("HOME"))
	configFile, err := toml.LoadFile(path)
	if err != nil {
		return nil, err
	}

	var config ConfigFile
	if err := configFile.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
