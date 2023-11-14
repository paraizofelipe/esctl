package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

const CONFIG_PATH = "%s/.config/esctl/config.toml"

type Cluster struct {
	Name     string   `toml:"name"`
	Address  []string `toml:"address"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
	Default  bool     `toml:"default"`
}

type Setup struct {
	Cluster []Cluster `toml:"cluster"`
}

func ReadSetup(filePath string) (setup Setup, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	if err = toml.Unmarshal(content, &setup); err != nil {
		return
	}

	return
}

func WriteSetup(setup Setup, filePath string) (err error) {
	res, err := toml.Marshal(setup)
	if err != nil {
		return
	}

	err = os.WriteFile(filePath, res, 0644)
	if err != nil {
		return
	}
	return
}

func (s *Setup) ClusterByName(name string) Cluster {
	for _, cluster := range s.Cluster {
		if cluster.Name == name {
			return cluster
		}
	}
	return Cluster{}
}

func (s *Setup) DefaultCluster() (cluster Cluster) {
	for _, cluster = range s.Cluster {
		if cluster.Default {
			return
		}
	}
	return
}
