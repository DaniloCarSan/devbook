package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

// Retorna um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routers.PublishInMux(r)
}
