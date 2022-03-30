package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// StringConnDB String de conexão com banco mysql
	StringConnDB = ""
	// Port - porta onde a API esta rodando
	Port = 0
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}
	StringConnDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))
}
