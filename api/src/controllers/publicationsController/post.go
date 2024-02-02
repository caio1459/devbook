package publicationscontroller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/caio1459/devbook/src/authentication"
	"github.com/caio1459/devbook/src/database"
	"github.com/caio1459/devbook/src/models"
	publicationrepositories "github.com/caio1459/devbook/src/repositories/publicationRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/caio1459/devbook/src/uploads"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
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
		return
	}
	publication.UserID = userID

	if err = publication.Prepare(); err != nil{
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
	publication.PubID, err = repositorie.InsertPublication(publication)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publication.ImageUrl != "" {
		if err = uploads.UploadHandler(publication.ImageUrl, publication.PubID, "publication"); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
	}
	responses.Json(w, http.StatusCreated, publication)
}
