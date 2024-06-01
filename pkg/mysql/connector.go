package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(config db.DB) (*connect, error) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s", db.DB.User, db.DB.Password, db.DB.Host, db.DB.Port, db.DB.Name)
	connect, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &connect, nil
}
