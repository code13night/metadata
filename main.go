package main

import (
	"fmt"

	"github.com/code13night/metadata/connector"
	"github.com/code13night/metadata/model"
)

func main() {
	cf := &model.Config{
		HostName:     "localhost",
		Port:         5432,
		UserName:     "postgres",
		Password:     "postgres",
		DatabaseName: "postgres",
		Options: map[string]string{
			"sslmode": "disable",
		},
	}
	connStr, errConn := connector.ConnPostgres(cf)
	if errConn != nil {
		fmt.Println(errConn)
		return
	}
	rsltMsg, errRslt := connector.CheckConnectionPostgres(connStr)
	if errRslt != nil {
		fmt.Println(errRslt)
		return
	}
	fmt.Println(rsltMsg)
}
