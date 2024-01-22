package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta a requisição em formato JSON
func Json(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatalf("Erro: %v", err)
		}
	}
}

// Retorna erro a requisição em formato JSON
func Erro(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
