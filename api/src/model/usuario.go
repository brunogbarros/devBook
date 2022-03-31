package model

import (
	"errors"
	"strings"
	"time"
)

// Usuario - modelo de um usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// validaUsuario : método para validar se os campos estão vazios
func (usuario *Usuario) validaUsuario() error {
	if usuario.Nome == "" {
		return errors.New("Nome é obrigatório!")
	}
	if usuario.Nick == "" {
		return errors.New("Nick é obrigatório!")
	}
	if usuario.Email == "" {
		return errors.New("Email é obrigatório!")
	}
	if usuario.Senha == "" {
		return errors.New("Senha é obrigatório!")
	}
	return nil
}

// formatar : formata strings para tirar espacos do inicio e do fim
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

// Preparar : chama os métodos para validar e formatar os dados recebidos do usuario
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validaUsuario(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}
