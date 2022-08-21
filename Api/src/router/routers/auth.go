package routers

import (
	"api/src/controllers"
	"net/http"
)

var routersAuth = RouteGroup{
	Name: "/auth",
	Routes: []Route{
		{
			Uri:                  "/sign/in",
			Method:               http.MethodPost,
			Exec:                 controllers.SignIn,
			RequireAuthenticated: false,
		},
	},
}
