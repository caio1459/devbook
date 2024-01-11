package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/caio1459/devbook/src/database"
	"github.com/caio1459/devbook/src/models"
	"github.com/caio1459/devbook/src/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Pega o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	//Converte o json em um struct
	user := models.User{}
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	//Abre uma conexão
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	//Cria um novo repositorio
	repositorie := repositories.NewRepositorieUser(db)
	userId, err := repositorie.Post(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("Usuário %v inserido com sucesso!", userId)))
}

func PostImage(w http.ResponseWriter, r *http.Request){
	
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando rota de get all"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando rota de get unico"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando rota de Put"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando rota de Delete"))
}
