package config

import (
	"flag"
	"os"
)

func Path() (path string) {
	var configPath string
	flag.StringVar(&configPath, "c", "", "Path to config")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("PROXY_GEO_CONFIG")
	}
	return configPath
}
