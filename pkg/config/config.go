package config

import (
	"os"
)

type Config struct {
	DBURL      string
	ServerAddr string
}

func Load() *Config {
	return &Config{
		DBURL:      getEnv("DB_URL", "postgres://postgres:super123@localhost:5432/postgres?sslmode=disable"),
		ServerAddr: getEnv("SERVER_ADDR", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
