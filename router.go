package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func appRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World")
	})

	buildHandler(router)

	return router
}
