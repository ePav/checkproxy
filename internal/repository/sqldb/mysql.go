package sqldb

import (
	"database/sql"
	"log"
)

func QueryDB(connect *sql.DB) (*sql.Rows, error) {
	rows, err := connect.Query("Select domain, ip, location From proxy")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("Error closing rows: %v", err)
		}
	}()

	return rows, nil
}
