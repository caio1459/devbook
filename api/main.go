package main

import (
	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caio1459/devbook/src/config"
	"github.com/caio1459/devbook/src/router"
)

//Gera uma chave para assinatura de token aleátoria para ser colocada no arquivo .env
// func init() {
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}
// 	stringKey := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringKey)
// }

func main() {
	config.Load()

	fmt.Printf("Escutando na porta %v", os.Getenv("PORT"))
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), r))
}
