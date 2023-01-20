package connector

import "github.com/code13night/metadata/model"

type ConnectionString interface {
	ConnSQLServer(cf *model.Config) (string, error)
	ConnPostgres(cf *model.Config) (string, error)
}
