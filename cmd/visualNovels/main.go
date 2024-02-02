package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n31t/golang-tsis/pkg/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/vn", handlers.WelcomeHandler).Methods("GET")
	r.HandleFunc("/vn/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/vn/all", handlers.PrintAllVisualNovels).Methods("GET")
	r.HandleFunc("/vn/{title}", handlers.GetVisualNovelByTitle).Methods("GET")
	http.ListenAndServe(":8080", r)
}
