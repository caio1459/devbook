package main

import (
	"log"
	"net/http"

	"github.com/caio1459/devbook/src/router"
)

func main() {
	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(":5000", r))
}
