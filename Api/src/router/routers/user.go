package routers

import (
	"api/src/controllers"
	"net/http"
)

var routersUser = RouteGroup{
	Name: "/user",
	Routes: []Route{
		{
			Uri:                  "",
			Method:               http.MethodPost,
			Exec:                 controllers.Create,
			RequireAuthenticated: false,
		},
		{
			Uri:                  "",
			Method:               http.MethodGet,
			Exec:                 controllers.All,
			RequireAuthenticated: false,
		},
		{
			Uri:                  "/{id}",
			Method:               http.MethodGet,
			Exec:                 controllers.ById,
			RequireAuthenticated: false,
		},
		{
			Uri:                  "/{id}",
			Method:               http.MethodPut,
			Exec:                 controllers.Update,
			RequireAuthenticated: false,
		},
		{
			Uri:                  "/{id}",
			Method:               http.MethodDelete,
			Exec:                 controllers.Delete,
			RequireAuthenticated: false,
		},
	},
}
