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

	repositorie := userrepositories.NewRepositorieUser(db)
	if err := repositorie.UpdateUser(uint64(ID), user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if user.ImageUrl != "" {
		//Faz o upload das imagens
		if err := uploads.UploadHandler(user.ImageUrl, IDToken, "user"); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
		}
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Usário %v atualizado com sucesso!", ID))
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
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel atualizar a senha de outro usuário"))
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

	if err = security.CheckPassword(password.Current, passwordDB); err != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("a senha atual está incorreta"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorie.UpdatePassword(ID, string(hashedPassword)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusCreated, fmt.Sprintf("Senha atualizada com sucesso!"))
}
