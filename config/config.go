package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	JwtSecret  string
	JwtExpired string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func LoadConfig() *Config {
	config := &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		JwtExpired: os.Getenv("JWT_EXPIRED"),
	}

	return config

}
