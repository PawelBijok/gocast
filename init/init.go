package initialization

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {

		log.Fatalln("Could not open env file")
	}

}
