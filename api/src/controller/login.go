package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repository"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login : Funcao do login para autenticar um usuario
func Login(w http.ResponseWriter, r *http.Request) {
	// fluxo:
	// ler o corpo da requisicão para pegar login e senha
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario model.Usuario
	if erro := json.Unmarshal(corpoRequisicao, usuario); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuario(db)
	usuarioSalvoNoBanco, erro := repo.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro = seguranca.VerificaSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	w.Write([]byte("Logado com sucesso. TODO: JWT"))
}
