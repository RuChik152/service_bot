package main

import (
	"belivr_service_bot/bot"
	"belivr_service_bot/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print("Success, .env file found")
	}

	db.MONGO_LOGIN, _ = os.LookupEnv("MONGO_LOGIN")
	if db.MONGO_LOGIN != "" {
		log.Println("MONGO_LOGIN:", db.MONGO_LOGIN)
	} else {
		log.Println("Ошибка!!! Не установлен логин к MongoDB")
		os.Exit(1)
	}
	db.MONGO_PASS, _ = os.LookupEnv("MONGO_PASS")
	if db.MONGO_PASS != "" {
		log.Println("MONGO_PASS:", "*************")
	} else {
		log.Println("Ошибка!!! Не установлен пароль к MongoDB")
		os.Exit(1)
	}
	db.MONGO_URL, _ = os.LookupEnv("MONGO_URL")
	if db.MONGO_URL != "" {
		log.Println("MONGO_URL:", db.MONGO_URL)
	} else {
		log.Println("Ошибка!!! Не установлен URI для подключения к MongoDB")
		os.Exit(1)
	}
	db.MONGO_DB_NAME, _ = os.LookupEnv("MONGO_DB_NAME")
	if db.MONGO_DB_NAME != "" {
		log.Println("MONGO_DB_NAME:", db.MONGO_DB_NAME)
	} else {
		log.Println("Ошибка!!! Не установлено имя БД в MongoDB")
		os.Exit(1)
	}

	db.MONGO_TYPE_CONNECT, _ = os.LookupEnv("MONGO_TYPE_CONNECT")
	if db.MONGO_TYPE_CONNECT != "" {
		log.Println("MONGO_TYPE_CONNECT:", db.MONGO_TYPE_CONNECT)
	} else {
		log.Println("Ошибка!!! Не установлено имя БД в MongoDB")
		os.Exit(1)
	}

	server = &http.Server{
		Addr:         ":9595",
		Handler:      appRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func main() {

	defer db.ConnectMongoDB()
	defer serverStop()

	go bot.InitBOT()

	go db.ConnectMongoDB()

	go startServer()

}

func startServer() {
	log.Printf("<<SERVER START>> http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func serverStop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	bot.BOT.StopReceivingUpdates()
	log.Println("Server stopped gracefully")
}
