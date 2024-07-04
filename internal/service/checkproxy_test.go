package service

import (
	"checkproxy/internal/repository/proxy"
	op "checkproxy/pkg/db"
	"database/sql"
	"fmt"
	"net"
	"os"
	"testing"
	"text/tabwriter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() (*sql.DB, error) {
	url := "testuser:testpassword@tcp(127.0.0.1:3306)/testdb"
	connect, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return connect, nil
}

func seedTestDB(connect *sql.DB) error {
	_, err := connect.Exec(`
		CREATE TABLE IF NOT EXISTS proxy (
			domain VARCHAR(255),
			ip VARCHAR(255),
			location VARCHAR(255)
		);
	`)
	if err != nil {
		return err
	}

	_, err = connect.Exec(`
		INSERT INTO proxy (domain, ip, location) VALUES
		('ProxyLine-us.pxl', '93.184.216.34', 'US'),
		('ProxyLine-tr.pxl', '93.184.216.34', 'TR');
	`)
	return err
}

func cleanupTestDB(connect *sql.DB) {
	connect.Exec("DROP TABLE IF EXISTS proxy;")
	connect.Close()
}

func TestCheckproxy(t *testing.T) {
	connect, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	defer cleanupTestDB(connect)

	err = seedTestDB(connect)
	if err != nil {
		t.Fatalf("Failed to seed test database: %v", err)
	}

	allproxies, err := proxy.QueryDB(connect)
	if err != nil {
		t.Fatalf("Error querying test database: %v", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.AlignRight|tabwriter.Debug)

	mmDB, err := op.Openmmdb()
	if err != nil {
		t.Fatalf("Error opening Maxmind database: %v", err)
	}
	ip2lDB, err := op.Openip2ldb()
	if err != nil {
		t.Fatalf("Error opening Ip2location database: %v", err)
	}

	for _, proxy := range allproxies {
		if proxy.Domain == "local" {
			continue
		}

		recordip2l, err := ip2lDB.Get_country_short(proxy.Ip)
		if err != nil {
			t.Fatalf("Error ip2location: %v", err)
		}

		recordmm, err := mmDB.Country(net.ParseIP(proxy.Ip))
		if err != nil {
			t.Fatalf("Error Maxmind: %v", err)
		}

		if assert.Equal(t, proxy.Location, recordip2l.Country_short) && assert.Equal(t, proxy.Location, recordmm.Country.IsoCode) {
			fmt.Printf("%s\t%s\t%s\tip2l OK, mm OK \n", proxy.Domain, proxy.Ip, proxy.Location)
		} else {
			fmt.Printf("%s\t%s\t%s\tError: ip2l %s, mm %s \n", proxy.Domain, proxy.Ip, proxy.Location, recordip2l.Country_short, recordmm.Country.IsoCode)
		}
		writer.Flush()
	}
}
