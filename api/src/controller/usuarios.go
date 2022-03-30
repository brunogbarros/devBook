package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CriarUsuario : create user
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var usuario model.Usuario
	// json -> struct
	if erro = json.Unmarshal(bodyReq, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()
	repositorie := repository.NovoRepositorioDeUsuario(db)
	usuarioId, erro := repositorie.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioId)))

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
