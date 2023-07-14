package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecret   string
	MongoDBURI  string
	MongoDBName string
	ServerHost  string
	ServerPort  string
}

var Cfg Config

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	Cfg = Config{
		JwtSecret:   os.Getenv("JWT_SECRET"),
		MongoDBURI:  os.Getenv("MONGODB_URI"),
		MongoDBName: os.Getenv("MONGODB_NAME"),
		ServerHost:  os.Getenv("SERVER_HOST"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}
}

func GetConfig() Config {
	return Cfg
}
