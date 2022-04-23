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
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
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
	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}
	return nil
}

// Deletar : delete by Id
func (repo usuariosRepository) Deletar(id uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}
	return nil
}

// BuscarPorEmail : busca um usuário por email e retorna o id e senha com hash
func (repo usuariosRepository) BuscarPorEmail(email string) (model.Usuario, error) {
	linha, erro := repo.db.Query("SELECT id, senha FROM usuarios WHERE email = ?")
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linha.Close()
	var usuario model.Usuario
	// percore as linhas, se existirem e popula o objeto usuario acima com os dados contidos na linha
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Seguir: funcao responsavel pela logica de seguir um usuario
func (repositorio usuariosRepository) Seguir(usuarioId uint, seguidorId uint) error {
	stmt, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil

}

// PararDeSeguir : funcao para parar de seguir um usuario
func (repositorio usuariosRepository) PararDeSeguir(usuarioId uint64, seguidorId uint64) error {
	stmt, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()
	if _, erro = stmt.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}
	return nil
}

// BuscarSeguidores : query de buscar seguidores na tabela
func (repositorio usuariosRepository) BuscarSeguidores(usuarioID uint64) ([]model.Usuario, error) {
	stmt, erro := repositorio.db.Query(`select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuarioId = ?`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer stmt.Close()

	var usuarios []model.Usuario
	for stmt.Next() {
		var usuario model.Usuario
		if erro = stmt.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Nick, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
