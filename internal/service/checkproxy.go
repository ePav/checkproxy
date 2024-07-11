package service

import (
	"fmt"
	"log"
	"net"

	"checkproxy/internal/repository/proxy"
	op "checkproxy/pkg/db"
	"os"
	"text/tabwriter"
)

func Checkproxy(allproxies []proxy.Proxy, config *op.Dbsource) {
	mmdb, err := op.Openmmdb(config.Database.Maxmind)
	if err != nil {
		log.Printf("Error opening Maxmind database: %v", err)
	}
	ip2ldb, err := op.Openip2ldb(config.Database.IP2Location)
	if err != nil {
		log.Printf("Error opening Ip2location database: %v", err)
	}

	defer mmdb.Close()
	defer ip2ldb.Close()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for _, proxy := range allproxies {
		if proxy.Domain == "local" {
			continue
		}

		recordip2l, err := ip2ldb.Get_country_short(proxy.IP)
		if err != nil {
			log.Printf("Error ip2location: %v", err)
		}

		recordmm, err := mmdb.Country(net.ParseIP(proxy.IP))
		if err != nil {
			log.Printf("Error Maxmind: %v", err)
		}

		if recordip2l.Country_short == proxy.Location && recordmm.Country.IsoCode == proxy.Location {
			fmt.Printf("%s\t%s\t%s\tip2l OK, mm OK \n", proxy.Domain, proxy.IP, proxy.Location)
		} else {
			fmt.Printf("%s\t%s\t%s\tError: ip2l %s, mm %s \n", proxy.Domain, proxy.IP, proxy.Location, recordip2l.Country_short, recordmm.Country.IsoCode)
		}
		writer.Flush()
	}
}
