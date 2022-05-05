package controller

import (
	"app/app/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("template/" +
	"*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ReturnDate()

	err := temp.ExecuteTemplate(w, "Index", products)
	if err != nil {
		return
	}
}

func Store(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "Create", nil)
	if err != nil {
		panic(err.Error())
	}
}
