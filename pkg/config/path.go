package config

import (
	"flag"
)

func Path() (Path string) {
	var configPath string
	flag.StringVar(&configPath, "c", "", "Path to config")
	flag.Parse()

	if configPath == "" {
		configPath = "internal/repository/config/config.yml"
	}
	return configPath
}
