package main

import (
	"app/app/routes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

const port = ":3000"

func main() {
	routes.Routes()

	fmt.Println("Server running in port:", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return
	}

}
