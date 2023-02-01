package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	value := os.Getenv(key)
	return value
}
