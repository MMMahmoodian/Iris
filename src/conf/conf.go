package conf

import (
	"github.com/joho/godotenv"
	"log"
)

func Initialize() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Loading .env failed. Err: %s", err)
	}
}
