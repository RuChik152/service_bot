package bot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var FTP_LOADER_URL string
var BUILD_SERVER_URL string

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
		log.Println("Ошибка!!! Не установлен URL к FTP LOADER")
		os.Exit(1)
	}

	BUILD_SERVER_URL, _ = os.LookupEnv("BUILD_SERVER_URL")
	if BUILD_SERVER_URL != "" {
		log.Println("BUILD_SERVER_URL: ", BUILD_SERVER_URL)
	} else {
		log.Println("Ошибка!!! Не установлен URL к BUILD_SERVER_URL")
		os.Exit(1)
	}

}
