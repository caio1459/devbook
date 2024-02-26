package routes

import (
	"net/http"

	commentcontroller "github.com/caio1459/devbook/src/controllers/commentController"
)

var routesComments = []Route{
	{
		URI:                    "/api/comments/{id}",
		Method:                 http.MethodPost,
		Function:               commentcontroller.CreateComment,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/api/comments/{id}",
		Method:                 http.MethodGet,
		Function:               commentcontroller.GetAllComments,
		RequiresAuthentication: true,
	},
}
