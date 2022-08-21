package routers

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Representa um objeto para as rotas da Api
type Route struct {
	Uri                  string
	Method               string
	Exec                 func(http.ResponseWriter, *http.Request)
	RequireAuthenticated bool
}

// Representa um grupo de rotas
type RouteGroup struct {
	Name   string
	Routes []Route
}

// Adiciona as rotas ao mux
func PublishInMux(r *mux.Router) *mux.Router {
	rgs := []RouteGroup{
		routersAuth,
		routersUser,
	}

	for _, rg := range rgs {
		for _, route := range rg.Routes {
			route.Uri = rg.Name + route.Uri

			if route.RequireAuthenticated {
				r.HandleFunc(route.Uri, middlewares.Authenticate(route.Exec)).Methods(route.Method)
			} else {
				r.HandleFunc(route.Uri, route.Exec).Methods(route.Method)
			}

		}
	}

	return r
}
