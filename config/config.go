package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	godotenv.Load()
	return os.Getenv(name)
}
