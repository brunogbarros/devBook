package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

//Buscar : Traz todos os usuários que existem com o nome ou nick inserido
func (repo usuariosRepository) Buscar(nomeOuNick string) ([]model.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // para usar o like = %nomeOuNick%
	linhas, erro := repo.db.Query("select id, nome,nick,email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []model.Usuario
	for linhas.Next() {
		var usuario model.Usuario
		// pega os dados da table ae joga na usuarios acima
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

// BuscarPorId : retorna um usuário por id
func (repo usuariosRepository) BuscarPorId(id uint64) (model.Usuario, error) {
	linhas, erro := repo.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", id)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linhas.Close()
	var usuario model.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}

//Atualizar : update user by Id
func (repo usuariosRepository) Atualizar(id uint64, usuario model.Usuario) error {
	statement, erro := repo.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email =? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}
	return nil
}
