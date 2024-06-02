package app

import (
	conf "checkproxy/pkg/config"
	"checkproxy/pkg/mysql"

	//"checkproxy/pkg/db"
	"checkproxy/internal/repository/sqldb"
	"log"
)

func Execute() {
	path := conf.Path()

	config, err := conf.LoadConfig(path)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connect, err := mysql.ConnectMySQL(*config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	rowsdb, err := sqldb.QueryDB(connect)
	if err != nil {
		log.Fatalf("Error queringing on database: %v", err)
	}

}
