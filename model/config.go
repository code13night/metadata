package model

type Config struct {
	HostName     string
	Port         int
	UserName     string
	Password     string
	DatabaseName string
	Options      map[string]string
}
