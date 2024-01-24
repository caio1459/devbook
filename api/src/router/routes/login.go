package routes

import (
	"net/http"

	"github.com/caio1459/devbook/src/controllers"
)

var routeLogin = Route{
	URI:                    "/api/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequiresAuthentication: true,
}
