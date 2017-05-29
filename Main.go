package main

import (
	"log"
	"net/http"
	"github.com/urfave/negroni"
)

func main() {

	router := NewRouter()

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}