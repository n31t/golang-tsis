package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type VisualNovel struct {
	Title       string `json:"title"`
	Developer   string `json:"developer"`
	ReleaseDate string `json:"release_date"`
	Platform    string `json:"platform"`
}

var data = []VisualNovel{
	VisualNovel{Title: "Fate/Stay Night", Developer: "Type-Moon", ReleaseDate: "January 30, 2004", Platform: "Windows"},
	VisualNovel{Title: "Fate/Hollow Ataraxia", Developer: "Type-Moon", ReleaseDate: "October 28, 2005", Platform: "Windows"},
}

func printAllVisualNovels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Visual Novels:\n")
	for _, v := range data {
		fmt.Fprint(w, v.Title, "\n")
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/vn", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Visual Novel Database!\n")
	})
	r.HandleFunc("/vn/all", printAllVisualNovels) // GET
	r.HandleFunc("/vn/{title}")
	http.ListenAndServe(":8080", r)
}
