package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/database"
	"github.com/caio1459/devbook/src/models"
	"github.com/caio1459/devbook/src/repositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/caio1459/devbook/src/uploads"
	"github.com/gorilla/mux"
)

func PostImage(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// Criar a estrutura de dados para armazenar os dados da imagem
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	image := models.Image{}

	//Converte o json para um struct
	if err = json.Unmarshal(requestBody, &image); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	//Abre uma conexão
	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieImage(db)

	image.ID, err = repositorie.PostImage(image, ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	uploads.UploadHandler(ID, image.FileName)
	responses.Json(w, http.StatusCreated, image)
}
