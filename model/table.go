package model

type Table struct {
	Name        string
	Owner       string
	Size        float64
	Columns     []string
	Constraints []string
	Triggers    []string
	Indexes     []string
	Definition  string
}
