package controller

import (
	"app/app/models"
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("template/" +
	"*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {
	products := models.ReturnDate()

	err := temp.ExecuteTemplate(w, "Index", products)
	if err != nil {
		return
	}
}

func Create(w http.ResponseWriter, _ *http.Request) {
	err := temp.ExecuteTemplate(w, "Create", nil)
	if err != nil {
		panic(err.Error())
	}
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	if r.RequestURI != "/store" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal("Parse form erro")
		return
	}

	values := r.PostForm

}
