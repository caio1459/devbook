package commentcontroller

import (
	"net/http"
	"strconv"

	"github.com/caio1459/devbook/src/database"
	commentrepositories "github.com/caio1459/devbook/src/repositories/commentRepositories"
	"github.com/caio1459/devbook/src/responses"
	"github.com/gorilla/mux"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	pubID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repositorie := commentrepositories.NewCommentRepositorie(db)
	comment, err := repositorie.SelectComments(pubID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusCreated, comment)
}
