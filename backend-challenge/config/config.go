package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppPort string

	ApiKey string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig(file string) *Config {
	err := godotenv.Load(file)
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		ApiKey:     getEnv("API_KEY", "apitest"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "order_food"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
