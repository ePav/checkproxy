package config

import (
	"checkproxy/pkg/db"
	"os"

	"gopkg.in/yaml.v2"
)

func loadConfig(configPath string) (*db.DB, error) {
	var config db.DB

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
