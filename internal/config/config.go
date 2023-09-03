package config

import (
	"github.com/pelletier/go-toml"
)

const CONFIG_PATH = "%s/.config/esctl/config.toml"

type Host struct {
	Name     string   `toml:"name"`
	Address  []string `toml:"address"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
	Default  bool     `toml:"default"`
}

type Setup struct {
	Host []Host `toml:"host"`
}

func Load(filePath string) (*Setup, error) {
	configFile, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	var setup Setup
	if err := configFile.Unmarshal(&setup); err != nil {
		return nil, err
	}

	return &setup, nil
}

func (s *Setup) HostByName(name string) Host {
	for _, host := range s.Host {
		if host.Name == name {
			return host
		}
	}
	return Host{}
}

func (s *Setup) DefaultHost() (host Host) {
	for _, host = range s.Host {
		if host.Default {
			return
		}
	}
	return
}
