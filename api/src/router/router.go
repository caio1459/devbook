package router

import (
	"github.com/caio1459/devbook/src/router/routes"
	"github.com/gorilla/mux"
)

// Cria um Router
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.RouteConfig(r)
}
