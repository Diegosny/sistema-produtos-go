package controller

import (
	"app/app/models"
	"html/template"
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

func Store(_ http.ResponseWriter, _ *http.Request) {

}
