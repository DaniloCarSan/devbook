package routers

import (
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

// Adicina as rotas ao mux
func PublishInMux(r *mux.Router) *mux.Router {
	rgs := []RouteGroup{
		routersUser,
	}

	for _, rg := range rgs {
		for _, route := range rg.Routes {
			route.Uri = rg.Name + route.Uri
			r.HandleFunc(route.Uri, route.Exec).Methods(route.Method)
		}
	}

	return r
}
