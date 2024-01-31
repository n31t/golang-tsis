package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n31t/golang-tsis/pkg/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/vn", handlers.WelcomeHandler)
	r.HandleFunc("/vn/health", handlers.HealthCheck)
	r.HandleFunc("/vn/all", handlers.PrintAllVisualNovels)
	r.HandleFunc("/vn/{title}", handlers.GetVisualNovelByTitle)
	http.ListenAndServe(":8080", r)
}
