package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// func LoadEnvVariables() {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	envPath := filepath.Join(filepath.Dir(ex), ".env")
// 	// fmt.Println(envPath)

// 	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
// 		err = godotenv.Load(envPath)
// 		if err != nil {
// 			log.Fatal("Error loading .env file")
// 		}
// 	}
// }

func LoadEnvVariables() {
    if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }
}
