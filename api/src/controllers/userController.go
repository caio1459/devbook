package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Acessando rota de post"))
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
