package db

import (
	"github.com/oschwald/geoip2-golang"
)

var countryMMDB string = "internal/repository/config/GeoLite2-Country.mmdb"

func Openmmdb() (*geoip2.Reader, error) {

	openmmdb, err := geoip2.Open(countryMMDB)
	if err != nil {
		return nil, err
	}
	return openmmdb, nil
}
