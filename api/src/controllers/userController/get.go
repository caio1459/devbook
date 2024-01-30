package usercontroller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/caio1459/devbook/src/database"
	userrepositories "github.com/caio1459/devbook/src/repositories/userRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//Pega o valor passado pela URL
	valueFilter := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := userrepositories.NewRepositorieUser(db)
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

	repositorie := userrepositories.NewRepositorieUser(db)
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

	repositorie := userrepositories.NewRepositorieUser(db)
	follows, err := repositorie.SelectFollows(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, follows)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
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

	repositorie := userrepositories.NewRepositorieUser(db)
	following, err := repositorie.SelectFollowing(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, following)
}
