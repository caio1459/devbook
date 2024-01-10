package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caio1459/devbook/src/config"
	"github.com/caio1459/devbook/src/router"
)

func main() {
	config.Load()

	fmt.Printf("Escutando na porta %v", os.Getenv("PORT"))
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), r))
}
