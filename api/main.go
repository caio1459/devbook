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
	"github.com/gorilla/handlers"
)

//*Gera uma chave para assinatura de token aleátoria para ser colocada no arquivo .env
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

	// Adicione o manipulador CORS aqui para autorização de outras aplicações
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"}) 

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), handlers.CORS(headers, methods, origins)(r)))
}
