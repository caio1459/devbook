package publicationscontroller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/authentication"
	"github.com/caio1459/devbook/src/database"
	publicationrepositories "github.com/caio1459/devbook/src/repositories/publicationRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/gorilla/mux"
)

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

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

	//Verifico se a publicação existe
	repositorie := publicationrepositories.NewRepositoriePublication(db)
	pubSavedDB, err := repositorie.SelectPublication(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if pubSavedDB.UserID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é permitido deletar a publicação de outro user"))
		return
	}

	if err = repositorie.DeletePublication(ID); err != nil{
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Publicação %v deletada com sucesso!", ID))
}
