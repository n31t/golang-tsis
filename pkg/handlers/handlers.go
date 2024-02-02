package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n31t/golang-tsis/pkg/models"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Visual Novel Database!\n")
}

func PrintAllVisualNovels(w http.ResponseWriter, r *http.Request) {
	visualNovels := models.GetAllVisualNovels()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(visualNovels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is healthy!\n")
}

func GetVisualNovelByTitle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	title := vars["title"]
	vn, err := models.GetVisualNovelByTitle(title)
	if err != true {
		template := template.Must(template.ParseFiles("ui/html/novn.html"))
		template.ExecuteTemplate(w, "novn.html", vn)
		return
	}

	template := template.Must(template.ParseFiles("ui/html/vn.html"))
	template.ExecuteTemplate(w, "vn.html", vn)
	// We can actually do the same stuff as in PrintAllVisualNovels, but we are using a template here
}
