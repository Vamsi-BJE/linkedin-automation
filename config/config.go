package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    string
	Password string
	MaxRuns  int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	return &Config{
		Email:    os.Getenv("LINKEDIN_EMAIL"),
		Password: os.Getenv("LINKEDIN_PASSWORD"),
		MaxRuns:  3,
	}
}
