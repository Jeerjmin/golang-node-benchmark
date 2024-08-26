package config

import (
	"log"
	"os"
)

type Config struct {
	Port               string
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDatabase string
}

func NewConfig() *Config {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	postgresqlHost := os.Getenv("POSTGRESQL_HOST")
	if postgresqlHost == "" {
		log.Fatal("POSTGRESQL_HOST is not set in .env file")
	}
	postgresqlPort := os.Getenv("POSTGRESQL_PORT")
	if postgresqlPort == "" {
		log.Fatal("POSTGRESQL_PORT is not set in .env file")
	}
	postgresqlUser := os.Getenv("POSTGRESQL_USER")
	if postgresqlUser == "" {
		log.Fatal("POSTGRESQL_USER is not set in .env file")
	}
	postgresqlPassword := os.Getenv("POSTGRESQL_PASSWORD")
	if postgresqlPassword == "" {
		log.Fatal("POSTGRESQL_PASSWORD is not set in .env file")
	}
	postgresqlDatabase := os.Getenv("POSTGRESQL_DATABASE")
	if postgresqlDatabase == "" {
		log.Fatal("POSTGRESQL_DATABASE is not set in .env file")
	}

	return &Config{
		Port:               port,
		PostgresqlHost:     postgresqlHost,
		PostgresqlPort:     postgresqlPort,
		PostgresqlUser:     postgresqlUser,
		PostgresqlPassword: postgresqlPassword,
		PostgresqlDatabase: postgresqlDatabase,
	}
}
