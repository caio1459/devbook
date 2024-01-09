package router

import "github.com/gorilla/mux"

//Cria um Router
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
