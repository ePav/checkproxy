package service

import (
	"fmt"
	"net"

	"checkproxy/internal/repository/proxy"
	"checkproxy/pkg/db"
	"os"
	"text/tabwriter"
)

func Checkproxy(proxies []proxy.Proxy, config *db.Dbsource) error {
	mmdb, err := db.Openmmdb(config.Database.Maxmind)
	if err != nil {
		return fmt.Errorf("error opening Maxmind database: %w", err)
	}
	ip2ldb, err := db.Openip2ldb(config.Database.IP2Location)
	if err != nil {
		return fmt.Errorf("error opening Ip2location database: %w", err)
	}

	defer mmdb.Close()
	defer ip2ldb.Close()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	for _, p := range proxies {
		if p.Domain == "local" {
			continue
		}

		recordip2l, err := ip2ldb.Get_country_short(p.IP)
		if err != nil {
			return fmt.Errorf("Error ip2location: %v", err)
		}

		recordmm, err := mmdb.Country(net.ParseIP(p.IP))
		if err != nil {
			return fmt.Errorf("Error Maxmind: %v", err)
		}

		if recordip2l.Country_short == p.Location && recordmm.Country.IsoCode == p.Location {
			fmt.Fprintf(writer, "%s\t%s\t%s\tip2l OK, mm OK\t\n", p.Domain, p.IP, p.Location)
		} else {
			fmt.Fprintf(writer, "%s\t%s\t%s\tError: ip2l %s, mm %s\t\n", p.Domain, p.IP, p.Location, recordip2l.Country_short, recordmm.Country.IsoCode)
		}
	}
	writer.Flush()

	return nil
}
