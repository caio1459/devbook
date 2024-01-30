package routes

import (
	"net/http"

	logincontroller "github.com/caio1459/devbook/src/controllers/loginController"
)

var routeLogin = Route{
	URI:                    "/api/login",
	Method:                 http.MethodPost,
	Function:               logincontroller.Login,
	RequiresAuthentication: false,
}
