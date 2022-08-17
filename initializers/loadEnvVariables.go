package initializers

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SECRET")
	fmt.Println(port)

}
