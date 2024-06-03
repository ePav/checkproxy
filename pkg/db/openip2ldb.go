package db

import (
	"github.com/ip2location/ip2location-go"
)

var countryIP2LDB string = "internal/repository/config/IP-COUNTRY-SAMPLE.BIN"

func Openip2ldb() (*ip2location.DB, error) {
	openip2ldb, err := ip2location.OpenDB(countryIP2LDB)
	if err != nil {
		return nil, err
	}
	return openip2ldb, nil
}
