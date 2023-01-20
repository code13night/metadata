package connector

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/code13night/metadata/model"
	_ "github.com/lib/pq"
)

func ConnPostgres(cf *model.Config) (string, error) {
	if cf == nil {
		return "", fmt.Errorf("configuration error")
	}
	cs := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cf.UserName, cf.Password, cf.HostName, strconv.Itoa(cf.Port), cf.DatabaseName, cf.Options["sslmode"])
	return cs, nil
}

func CheckConnectionPostgres(connStr string) (string, error) {
	db, err := sql.Open("postgres", connStr)
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

func GetDatabases(db *sql.DB) ([]model.Database, error) {
	var dbms []model.Database
	var dbm model.Database
	rows, err := db.Query("SELECT * FROM album")

	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("There is an issue getting the data")
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&dbm.Name, &dbm.Owner, &dbm.Tablespace, &dbm.Collation, &dbm.CreatedAt, &dbm.Size, &dbm.Definition)
		dbms = append(dbms, dbm)
	}
	return dbms, nil
}
