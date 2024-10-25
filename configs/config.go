package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB DbConfig
}
type DbConfig struct {
	Dsn string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DB: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
