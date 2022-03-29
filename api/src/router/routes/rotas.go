package rotas

import "net/http"

// Rota : define a estrutura basica de uma rota da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
