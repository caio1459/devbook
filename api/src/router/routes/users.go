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
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/follow",
		Method:                 http.MethodDelete,
		Function:               controllers.StopFollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/follows",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollows,
		RequiresAuthentication: true,
	},
}
