package main

import (
	"belivr_service_bot/bot"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func buildHandler(router *mux.Router) {
	buildRouter := router.PathPrefix("/building").Subrouter()

	buildRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>building</h1>")
	})

	buildRouter.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		log.Println("<<<<ПОЛУЧИЛ ЗАПРОС ОТ СЕРВЕРА>>>>")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		log.Println("<<<<ПОЛУЧИЛ ЗАПРОС ОТ СЕРВЕРА>>>>", string(body))
		bot.BOT_CHANEL <- body
		log.Println("Received data: ", string(body))
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

}
