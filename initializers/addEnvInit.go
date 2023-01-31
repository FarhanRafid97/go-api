package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializerEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}
