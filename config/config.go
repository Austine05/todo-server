package config

import (
	"log"

	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port        string
	JWTSecret   string
	MongoDBURI  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		MongoDBURI:  os.Getenv("MONGODB_URI"),
	}
}
