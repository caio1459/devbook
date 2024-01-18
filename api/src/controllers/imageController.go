package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/database"
	"github.com/caio1459/devbook/src/models"
	"github.com/caio1459/devbook/src/repositories"
	"github.com/caio1459/devbook/src/uploads"
	"github.com/gorilla/mux"
)

func PostImage(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		log.Fatalf("Erro ao converter o parâmetro para um inteiro: %v", err)
	}

	// Criar a estrutura de dados para armazenar os dados da imagem

	// Decodificar o corpo da solicitação JSON e armazenar na estrutura de dados
	//decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&image); err != nil {
	// 	log.Fatal(err)
	// }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	image := models.Image{}

	if err = json.Unmarshal(body, &image); err != nil {
		log.Fatal(err)
	}

	//Abre uma conexão
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	repositorie := repositories.NewRepositorieImage(db)

	Id, err := repositorie.PostImage(image, ID)
	if err != nil {
		log.Fatal(err)
	}

	uploads.UploadHandler(ID, image.FileName)

	// Responder com sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Upload de imagem cadastrada com sucesso para o usuario %v\n", Id)
}
