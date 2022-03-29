package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Rota : define a estrutura basica de uma rota da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar : returns a configured router with all routes
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
