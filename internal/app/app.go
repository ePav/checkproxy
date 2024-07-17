package app

import (
	conf "checkproxy/pkg/config"
	"checkproxy/pkg/mysql"

	"checkproxy/internal/repository/proxy"
	"checkproxy/internal/service"

	"log"
)

func Execute() {
	path := conf.Path()

	config, err := conf.LoadConfig(path)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connect, err := mysql.ConnectMySQL(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	proxies, err := proxy.GetAll(connect)
	if err != nil {
		log.Fatalf("Error queringing on database: %v", err)
	}

	if err := service.Checkproxy(proxies, config); err != nil {
		log.Fatalf("Error checking proxy: %v", err)
	}
}
