package model

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
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
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// validar : valida se os campos não estão em branco
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode ficar em branco")
	}
	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatório e não pode ficar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é obrigatório e não pode ficar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode ficar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("E-mail inserido é inválido.")
	}
	return nil
}

// formatar : tira os espacos da strings
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
