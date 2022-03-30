package repository

import (
	"api/src/model"
	"database/sql"
)

// O repositori vai interagir com o banco
// Esse struct terá todos os métodos necessário para o CRUD do banco
type usuariosRepository struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario retorna um ponteiro para o struct usuariosRepository
func NovoRepositorioDeUsuario(db *sql.DB) *usuariosRepository {
	// passamos o end do struct acima passando a instancia do db
	return &usuariosRepository{db}
}

// Criar : método para inserir no banco de dados
func (repo usuariosRepository) Criar(usuario model.Usuario) (uint64, error) {
	statement, erro := repo.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	// passando do Exec o usuário já foi adicionado no banco, basta só pegar o id retornado
	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIdInserido), nil
}
