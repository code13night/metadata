package model

type Metadata struct {
	Database     string
	Schema       string
	Table        string
	Password     string
	DatabaseName string
	Options      map[string]string
}
