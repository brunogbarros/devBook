package controller

import "net/http"

//CriarUsuario : create user
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

//BuscarUsuarios: get all users
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuário"))
}

//BuscaUsuario : get user by id
func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando 1 usuário"))
}

//AtualizarUsuario : update user
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeletarUsuario : delete user
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuário"))
}
