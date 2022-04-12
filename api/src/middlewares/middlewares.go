package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger : um exemplo de logger para printar informacoes no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar : middleware para autenticacão
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando token... todo")
		next(w, r)
	}
}
