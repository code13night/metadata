package model

type Column struct {
	Name         string
	DataType     string
	Length       int64
	Scale        int64
	DefaultValue string
	IsNull       bool
	IsPrimary    bool
	IsForeign    bool
	Constraints  []string
	Indexes      []string
}
