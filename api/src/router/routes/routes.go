package routes

import "net/http"

//Representa todas as rotas da API
type Route struct {
	URI                    string                                       //Uri da rota
	Method                 string                                       //Método da rota
	Function               func(w http.ResponseWriter, r *http.Request) //Funcão da Rota
	RequiresAuthentication bool                                         //Requer auticação de rota
}
