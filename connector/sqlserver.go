package connector

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/code13night/metadata/model"
	_ "github.com/microsoft/go-mssqldb"
)

func ConnSQLServer(cf *model.Config) (string, error) {
	if cf != nil {
		return "", fmt.Errorf("configuration error")
	}
	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(cf.UserName, cf.Password),
		Host:   fmt.Sprintf("%s:%d", cf.HostName, cf.Port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	return u.String(), nil
}

func CheckConnection(dbType string, connStr string) (string, error) {
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return "", fmt.Errorf("error connecting to the database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing the database connection: %v", err)
		}
	}()

	pingErr := db.Ping()
	if pingErr != nil {
		return "", pingErr
	}
	return "Connected", nil
}
