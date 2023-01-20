package model

type Database struct {
	Name       string
	Owner      string
	Tablespace string
	Collation  string
	CreatedAt  string
	Size       float64
	Schemas    []string
	Options    map[string]string
	Definition string
}
