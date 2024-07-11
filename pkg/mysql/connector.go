package mysql

import (
	"checkproxy/pkg/db"
	"database/sql"
	"fmt"

	// Importing the MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(config *db.Dbsource) (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	connect, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return connect, nil
}
