package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Representa todas as rotas da API
type Route struct {
	URI                    string                                       //Uri da rota
	Method                 string                                       //Método da rota
	Function               func(w http.ResponseWriter, r *http.Request) //Funcão da Rota
	RequiresAuthentication bool                                         //Requer auticação de rota
}

// Coloca todas as rotas no router
func RouteConfig(r *mux.Router) *mux.Router {
	routes := routesUsers
	//Adiciona mais rotas
	routes = append(routes, routeLogin)
	//Separa as rotas existentes
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
