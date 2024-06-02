package mysql

import (
	"checkproxy/pkg/db"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(config db.DB) (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
	connect, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return connect, nil
}
