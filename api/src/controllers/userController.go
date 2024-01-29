package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/caio1459/devbook/src/authentication"
	"github.com/caio1459/devbook/src/database"
	"github.com/caio1459/devbook/src/models"
	"github.com/caio1459/devbook/src/repositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/caio1459/devbook/src/uploads"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Pega o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Converte o json em um struct
	user := models.User{}
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Verifica os campos
	if err = user.Prepare("cadastro"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Abre uma conexão
	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	//Cria um novo repositorio
	repositorie := repositories.NewRepositorieUser(db)
	user.ID, err = repositorie.PostUser(user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if user.ImageUrl != "" {
		//Faz o upload das imagens
		if err := uploads.UploadHandler(user.ImageUrl); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
		}
	}
	responses.Json(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//Pega o valor passado pela URL
	valueFilter := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	users, err := repositorie.SelectUsers(valueFilter)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, err := strconv.ParseUint(parameter["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	user, err := repositorie.SelectUser(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	//Verifica se existe o usuário
	if user.Register.IsZero() {
		responses.Json(w, http.StatusBadRequest, fmt.Sprintf("Usário com id %v não encontrado", ID))
		return
	}
	responses.Json(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, err := strconv.ParseInt(parameter["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	IDToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Só deixa o usuário atualizar a si mesmo
	if ID != int64(IDToken) {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel alterar esse usuário"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = user.Prepare("atualizar"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	if err := repositorie.UpdateUser(uint64(ID), user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if user.ImageUrl != "" {
		//Faz o upload das imagens
		if err := uploads.UploadHandler(user.ImageUrl); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
		}
	}
	responses.Json(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	ID, err := strconv.ParseInt(parameter["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	IDToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Só deixa o usuário atualizar a si mesmo
	if ID != int64(IDToken) {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel deletar esse usuário"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	err = repositorie.DeleteUser(uint64(ID))
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	followID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if followID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel seguir você mesmo"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	if err = repositorie.PostFollow(userID, followID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)
}

func StopFollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	followID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if followID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel deixar de seguir você mesmo"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	if err = repositorie.DeleteFollow(userID, followID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)
}

func GetFollows(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewRepositorieUser(db)
	follows, err := repositorie.SelectFollows(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, follows)
}
