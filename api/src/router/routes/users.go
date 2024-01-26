package routes

import (
	"net/http"

	"github.com/caio1459/devbook/src/controllers"
)

// Cria rotas de usuarios
var routesUsers = []Route{
	{
		URI:                    "/api/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/api/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetAllUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/api/users/upload/{id}",
		Method:                 http.MethodPost,
		Function:               controllers.PostImage,
		RequiresAuthentication: false,
	},
}
