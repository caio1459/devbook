package usercontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/authentication"
	"github.com/caio1459/devbook/src/database"
	"github.com/caio1459/devbook/src/models"
	userrepositories "github.com/caio1459/devbook/src/repositories/userRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/caio1459/devbook/src/security"
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
	repositorie := userrepositories.NewRepositorieUser(db)
	user.ID, err = repositorie.InsertUser(user)
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

	repositorie := userrepositories.NewRepositorieUser(db)
	if err = repositorie.InsertFollow(userID, followID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Usário %v seguido com sucesso!", followID))
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if ID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não atualizar a senha de outro usuário"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	password := models.Password{}
	if err = json.Unmarshal(requestBody, &password); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := userrepositories.NewRepositorieUser(db)
	passwordDB, err := repositorie.SelectUserPassword(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(password.Current, passwordDB); err != nil{
		responses.Erro(w, http.StatusUnauthorized, errors.New("a senha atual está incorreta"))
		return
	}
}
