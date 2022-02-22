package main

import (
	"fmt"
	"net/http"

	"./views"

	"github.com/gorilla/mux"
)

var (
	homeTpl    *views.View
	aboutusTpl *views.View
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTpl.Template.ExecuteTemplate(w, homeTpl.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func aboutusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := aboutusTpl.Template.ExecuteTemplate(w, aboutusTpl.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
}

func main() {
	homeTpl = views.NewView("bootstrap", "views/home.gohtml")
	aboutusTpl = views.NewView("bootstrap", "views/aboutus.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/aboutus", aboutusHandler)
	http.ListenAndServe(":3000", r)
}
