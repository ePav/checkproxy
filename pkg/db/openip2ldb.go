package db

import (
	"github.com/ip2location/ip2location-go"
)

func Openip2ldb(countryIP2LDB string) (*ip2location.DB, error) {
	openip2ldb, err := ip2location.OpenDB(countryIP2LDB)
	if err != nil {
		return nil, err
	}
	return openip2ldb, nil
}
