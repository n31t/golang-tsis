package handlers

import (
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
	template := template.Must(template.ParseFiles("ui/html/allvn.html"))
	template.ExecuteTemplate(w, "allvn.html", struct{ VisualNovels []models.VisualNovel }{models.GetAllVisualNovels()})

	// fmt.Fprintf(w, "All Visual Novels:\n")
	// for _, v := range models.GetAllVisualNovels() {
	// 	fmt.Fprint(w, v.Title, "\n")
	// }
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
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "Visual Novel not found!\n")
		template := template.Must(template.ParseFiles("ui/html/novn.html"))
		template.ExecuteTemplate(w, "novn.html", vn)
		return
	}

	template := template.Must(template.ParseFiles("ui/html/vn.html"))
	template.ExecuteTemplate(w, "vn.html", vn)
	// fmt.Fprintf(w, "Visual Novel Details:\nTitle: %s\nDeveloper: %s\nRelease Date: %s\nPlatform: %s\n",
	// 	vn.Title, vn.Developer, vn.ReleaseDate, vn.Platform)
}
