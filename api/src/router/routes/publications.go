package routes

import (
	"net/http"

	publicationscontroller "github.com/caio1459/devbook/src/controllers/publicationsController"
)

var routesPublications = []Route{
	{
		URI:                    "/api/publications",
		Method:                 http.MethodPost,
		Function:               publicationscontroller.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/publications",
		Method:                 http.MethodGet,
		Function:               publicationscontroller.GetAllPublications,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/publication/{id}",
		Method:                 http.MethodGet,
		Function:               publicationscontroller.GetPublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/publication/{id}",
		Method:                 http.MethodPut,
		Function:               publicationscontroller.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/publication/{id}",
		Method:                 http.MethodDelete,
		Function:               publicationscontroller.DeletePublication,
		RequiresAuthentication: true,
	},
}
