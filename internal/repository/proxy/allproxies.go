package proxy

import (
	"database/sql"
	"log"
)

type Proxy struct {
	Domain   string
	IP       string
	Location string
}

func QueryDB(connect *sql.DB) ([]Proxy, error) {
	rows, err := connect.Query("SELECT domain, ip, location FROM proxy")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("Error closing rows: %v", err)
		}
	}()

	var proxies []Proxy
	for rows.Next() {
		var proxy Proxy
		err := rows.Scan(&proxy.Domain, &proxy.IP, &proxy.Location)
		if err != nil {
			return nil, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, nil
}
