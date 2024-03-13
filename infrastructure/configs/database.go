package configs

import (
	"log"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     int
	Dialect  string
}

func GetDatabaseConfig() *DatabaseConfig {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	if err != nil {
		log.Fatal("Invalid Port")
	}

	return &DatabaseConfig{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Name:     os.Getenv("DATABASE_NAME"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		Dialect:  os.Getenv("DATABASE_DIALECT"),
	}
}
