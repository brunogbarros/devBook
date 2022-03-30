package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver Mysql importado 'na mão'
	"log"
)

// Conectar - Abre a conexão com o banco e retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnDB)
	if erro != nil {
		log.Fatal("Não conseguiu abrir o banco!")
	}

	if erro = db.Ping(); erro != nil {
		// se deu erro fecha
		_ = db.Close()
		return nil, erro
	}
	return db, nil
}
