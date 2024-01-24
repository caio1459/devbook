package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnection string //String de conexão com o banco
	Port             int    //Porta do servidor
	err              error  //Erro de retorno de funções
	SecretKey        []byte //Chave para assinar o token
)

// Inicia as variaveis de ambiente
func Load() {
	if err = godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar variaveis de ambiente: %v", err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	//Cria uma string de conexão com base nas variaveis de ambiente do arquivo.env
	StringConnection = fmt.Sprintf(
		"%v:%v@/%v?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
