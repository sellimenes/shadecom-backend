package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
    if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }
}
