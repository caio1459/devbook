package usercontroller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/authentication"
	"github.com/caio1459/devbook/src/database"
	userrepositories "github.com/caio1459/devbook/src/repositories/userRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/gorilla/mux"
)

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

	repositorie := userrepositories.NewRepositorieUser(db)
	err = repositorie.DeleteUser(uint64(ID))
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Usário %v deletado com sucesso!", ID))
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

	repositorie := userrepositories.NewRepositorieUser(db)
	if err = repositorie.DeleteFollow(userID, followID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Usário %v deixado de seguir com sucesso!", followID))
}
