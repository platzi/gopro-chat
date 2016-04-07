package lib

import (
	"fmt"
	"upper.io/db"
	"upper.io/db/postgresql"
)

type Config struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Name:     "platzi",
		User:     "platzi",
		Password: "platzi",
		Port:     5432,
	}
}

func (c *Config) ConnectionString() (conn postgresql.ConnectionURL) {
	return postgresql.ConnectionURL{
		User:     c.User,
		Address:  db.HostPort(c.Host, uint(c.Port)),
		Database: c.Name,
		Password: c.Password,
	}
}

func (c *Config) String() (conn string) {
	conn = fmt.Sprintf("host=%s dbname=%s sslmode=disable user=%s password=%s", c.Host, c.Name, c.User, c.Password)
	return
}
