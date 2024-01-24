package middlewares

import (
	"log"
	"net/http"
	"os"
)

// Verifica se o user fazendo a requisição está autenticado
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Validando..."))
		next(w, r) //Executa próxima função
	}
}

//Gera logs de rotas
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Abrir ou criar um arquivo de logs
		file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("Erro ao abrir o arquivo de logs:", err)
		}
		defer file.Close()
		// Configurar o logger para escrever no arquivo
		log.SetOutput(file)
		log.Printf(" | Método: %s | URI: %s", r.Method, r.RequestURI)

		next(w, r)
	}
}
