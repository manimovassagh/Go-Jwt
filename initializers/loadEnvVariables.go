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

		fmt.Println("some")
	}

	port := os.Getenv("SECRET")
	fmt.Println(port)

}
