package routes

import (
	"app/app/controller"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/create", controller.Store)
}
