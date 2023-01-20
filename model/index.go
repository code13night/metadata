package model

type Index struct {
	Name       string
	Owner      string
	Method     string
	IsUnique   bool
	IsCluster  bool
	Columns    []string
	Definition string
}
