package publicationscontroller

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
	publicationrepositories "github.com/caio1459/devbook/src/repositories/publicationRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/caio1459/devbook/src/uploads"
	"github.com/gorilla/mux"
)

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
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
		responses.Erro(w, http.StatusForbidden, errors.New("não é permitido atualizar a publicação de outro user"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	publication := models.Publications{}
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}

	if err = publication.Prepare(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorie.UpdatePublication(ID, publication); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publication.ImageUrl != "" {
		if err = uploads.UploadHandler(publication.ImageUrl, publication.PubID, "publication"); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Publicação %v atualizada com sucesso!", ID))
}

func Likes(w http.ResponseWriter, r *http.Request) {
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
	if err = repositorie.UpdateLikes(ID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Publicação %v curtida", ID))
}

func Deslikes(w http.ResponseWriter, r *http.Request) {
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
	if err = repositorie.UpdateDeslikes(ID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, fmt.Sprintf("Publicação %v Descurtida", ID))
}
