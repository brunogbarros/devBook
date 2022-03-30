package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON transforma os dados recebidos em JSON ----------\/---------ou uma interface{}
func JSON(w http.ResponseWriter, statusCode int, dados any) {
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{Erro: erro.Error()})
}
