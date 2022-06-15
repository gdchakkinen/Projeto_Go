package main

import (
	"net/http"

	"go_modules/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
