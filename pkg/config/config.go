package config

import (
	"checkproxy/pkg/db"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configPath string) (*db.Dbsource, error) {
	var config db.Dbsource

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
