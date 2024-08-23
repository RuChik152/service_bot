package bot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var FTP_LOADER_URL string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print("Success, .env file found")
	}

	FTP_LOADER_URL, _ = os.LookupEnv("FTP_LOADER_URL")
	if FTP_LOADER_URL != "" {
		log.Println("FTP_LOADER_URL", FTP_LOADER_URL)
	} else {
		log.Println("Ошибка!!! Не установлен логин к MongoDB")
		os.Exit(1)
	}

}
