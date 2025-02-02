package test

import (
	"checkproxy/internal/repository/proxy"
	op "checkproxy/pkg/db"
	"database/sql"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"text/tabwriter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func GetProjectPath() (string, error) {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to obtain project path")
	}

	root := filepath.Join(filepath.Dir(b), "..", "..")
	return root, nil
}

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
		('ProxyLine-tr.pxl', '38.180.112.170', 'TR');
	`)
	return err
}

func cleanupTestDB(connect *sql.DB) {
	_, err := connect.Exec("DROP TABLE IF EXISTS proxy;")
	if err != nil {
		fmt.Printf("Failed to drop table: %v", err)
	}
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

	pathProject, err := GetProjectPath()
	if err != nil {
		t.Fatalf("Failed to get project path: %v", err)
	}
	ip2lDB, err := op.Openip2ldb(pathProject + "/internal/repository/config/IP-COUNTRY-SAMPLE.BIN")
	if err != nil {
		t.Fatalf("Error opening Ip2location database: %v", err)
	}
	mmDB, err := op.Openmmdb(pathProject + "/internal/repository/config/GeoLite2-Country.mmdb")
	if err != nil {
		t.Fatalf("Error opening Maxmind database: %v", err)
	}
	defer mmDB.Close()
	defer ip2lDB.Close()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for _, proxy := range allproxies {
		if proxy.Domain == "local" {
			continue
		}

		recordip2l, err := ip2lDB.Get_country_short(proxy.IP)
		if err != nil {
			t.Fatalf("Error ip2location: %v", err)
		}

		recordmm, err := mmDB.Country(net.ParseIP(proxy.IP))
		if err != nil {
			t.Fatalf("Error Maxmind: %v", err)
		}

		if assert.Equal(t, proxy.Location, recordip2l.Country_short) && assert.Equal(t, proxy.Location, recordmm.Country.IsoCode) {
			fmt.Printf("%s\t%s\t%s\tip2l OK, mm OK \n", proxy.Domain, proxy.IP, proxy.Location)
		} else {
			fmt.Printf("%s\t%s\t%s\tError: ip2l %s, mm %s \n", proxy.Domain, proxy.IP, proxy.Location, recordip2l.Country_short, recordmm.Country.IsoCode)
		}
		writer.Flush()
	}
}
