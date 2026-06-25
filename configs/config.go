package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB string
}

func Load() Config {

	godotenv.Load()

	return Config{
		DB: os.Getenv("DATABASE_URL"),
	}
}
