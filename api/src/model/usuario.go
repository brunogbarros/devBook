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

// Preparar : método para preparar e formatar o usuário no padrão
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

// validar : valida se os campos não estão em branco
func (usuario *Usuario) validar() error {
	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode ficar em branco")
	}
	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatório e não pode ficar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("A Senha é obrigatório e não pode ficar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode ficar em branco")
	}
	return nil
}

// formatar : tira os espacos da strings
func (usuario *Usuario) formatar() {
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
}
