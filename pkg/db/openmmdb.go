package db

import (
	"github.com/oschwald/geoip2-golang"
)

func Openmmdb(countryMMDB string) (*geoip2.Reader, error) {
	openmmdb, err := geoip2.Open(countryMMDB)
	if err != nil {
		return nil, err
	}
	return openmmdb, nil
}
