package config

import "os"

type Config struct {
	DatabaseURL string
}

func LoadConfig() Config {
	URL := os.Getenv("DATABASE_URL")

	if URL == "" {
		URL = "postgres://postgres:root@localhost:5432/sydesk?sslmode=disable"
	}

	return Config{
		DatabaseURL: URL,
	}
}
