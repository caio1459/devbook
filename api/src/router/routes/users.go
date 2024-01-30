package routes

import (
	"net/http"

	usercontroller "github.com/caio1459/devbook/src/controllers/userController"
)

// Cria rotas de usuarios
var routesUsers = []Route{
	{
		URI:                    "/api/users",
		Method:                 http.MethodPost,
		Function:               usercontroller.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/api/users",
		Method:                 http.MethodGet,
		Function:               usercontroller.GetAllUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodGet,
		Function:               usercontroller.GetUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodPut,
		Function:               usercontroller.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}",
		Method:                 http.MethodDelete,
		Function:               usercontroller.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/follow",
		Method:                 http.MethodPost,
		Function:               usercontroller.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/follow",
		Method:                 http.MethodDelete,
		Function:               usercontroller.StopFollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/follows",
		Method:                 http.MethodGet,
		Function:               usercontroller.GetFollows,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/following",
		Method:                 http.MethodGet,
		Function:               usercontroller.GetFollowing,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/users/{id}/update-password",
		Method:                 http.MethodPost,
		Function:               usercontroller.UpdatePassword,
		RequiresAuthentication: true,
	},
}
