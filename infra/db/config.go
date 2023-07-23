package db

import "fmt"

type Config struct {
	Username   string `json:"username" env:"POSTGRES_USERNAME" required:"true"`
	Password   string `json:"password" env:"POSTGRES_PASSWORD" required:"true"`
	Host       string `json:"host" env:"POSTGRES_HOST" required:"true"`
	Port       string `json:"port" env:"POSTGRES_PORT" required:"true"`
	Database   string `json:"database" env:"POSTGRES_DATABASE" required:"true"`
	Connection string `json:"connection" env:"POSTGRES_CONNECTION" required:"true"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.Username, c.Password, c.Database, c.Port,
	)
}
