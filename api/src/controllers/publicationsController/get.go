package publicationscontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/authentication"
	"github.com/caio1459/devbook/src/database"
	publicationrepositories "github.com/caio1459/devbook/src/repositories/publicationRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/gorilla/mux"
)

func GetAllPublications(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := publicationrepositories.NewRepositoriePublication(db)
	publications, err := repositorie.SelectPublications(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, publications)
}

func GetPublication(w http.ResponseWriter, r *http.Request) {
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

	repositorie := publicationrepositories.NewRepositoriePublication(db)
	publication, err := repositorie.SelectPublication(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	//Verifica se existe o usuário
	if publication.Register.IsZero() {
		responses.Json(w, http.StatusBadRequest, fmt.Sprintf("Publicação com id %v não encontrado", ID))
		return
	}
	responses.Json(w, http.StatusOK, publication)
}

func GetPublicationsFromUser(w http.ResponseWriter, r *http.Request) {
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

	repositorie := publicationrepositories.NewRepositoriePublication(db)
	publications, err := repositorie.SelectPublicationsFromUser(ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, publications)
}
