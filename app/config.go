package main

import "os"

type Config struct {
	DatabaseURL string
	BaseURL     string
	Port        string
}

func LoadConfig() Config {
	return Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://..."),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
		Port:        getEnv("PORT", ":8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
