package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppPort           string
	DbHost            string
	DbPort            string
	DbUser            string
	DbPassword        string
	DbName            string
	JwtSecret         string
	JwtAccessExpired  int
	JwtRefreshExpired int
}

var Data *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	JwtAccessExpiration := os.Getenv("JWT_ACCESS_EXPIRATION_MINUTES")
	if JwtAccessExpiration == "" {
		log.Fatalf("Environment variable JWT_ACCESS_EXPIRATION_MINUTES is not set")
	}
	JwtAccessExpirationVar, err := strconv.Atoi(JwtAccessExpiration)
	if err != nil {
		log.Fatalf("Error converting environment variable to integer: %v\n", err)
	}

	JwtRefreshExpiration := os.Getenv("JWT_ACCESS_EXPIRATION_MINUTES")
	if JwtRefreshExpiration == "" {
		log.Fatalf("Environment variable JWT_ACCESS_EXPIRATION_MINUTES is not set")
	}
	JwtRefreshExpirationVar, err := strconv.Atoi(JwtAccessExpiration)
	if err != nil {
		log.Fatalf("Error converting environment variable to integer: %v\n", err)
	}

	config := &Config{
		AppPort:           os.Getenv("APP_PORT"),
		DbHost:            os.Getenv("DB_HOST"),
		DbPort:            os.Getenv("DB_PORT"),
		DbUser:            os.Getenv("DB_USER"),
		DbPassword:        os.Getenv("DB_PASSWORD"),
		DbName:            os.Getenv("DB_NAME"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		JwtAccessExpired:  JwtAccessExpirationVar,
		JwtRefreshExpired: JwtRefreshExpirationVar,
	}
	Data = config
}
