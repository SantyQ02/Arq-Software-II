package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// func init() {
// 	LoadEnv()
// }

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func LoadTestEnv(env string) {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}
