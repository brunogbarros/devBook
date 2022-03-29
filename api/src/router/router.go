package router

import (
	rotas "api/src/router/routes"
	"github.com/gorilla/mux"
)

//GetRouter : returns a router with all routes
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
