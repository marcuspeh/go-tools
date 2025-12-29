package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Init(envFile string) {
	if envFile == "" {
		envFile = ".env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		panic("Error loading .env file")
	}
}

func GetEnvString(key string) string {
	return os.Getenv(key)
}
