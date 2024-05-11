package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa as informações de um usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(operacao string) error {
	if erro := usuario.validar(operacao); erro != nil {
		return erro
	}

	if erro := usuario.formatar(operacao); erro != nil {
		return erro
	}
	return nil
}

// validar realiza a validação dos campos do usuário com base no tipo de operação
func (usuario *Usuario) validar(operacao string) error {
	if len(usuario.Nome) < 3 || len(usuario.Nome) > 50 {
		return errors.New("o nome deve ter entre 3 e 50 caracteres")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if len(usuario.Nick) < 3 || len(usuario.Nick) > 20 {
		return errors.New("o nick deve ter entre 3 e 20 caracteres")
	}

	// Validação condicional da senha baseada no tipo de operação
	if operacao == "cadastro" {
		if len(usuario.Senha) < 6 || len(usuario.Senha) > 30 {
			return errors.New("a senha deve ter entre 6 e 30 caracteres")
		}

	} else if operacao == "edicao" && len(usuario.Senha) > 0 {
		if len(usuario.Senha) < 6 || len(usuario.Senha) > 30 {
			return errors.New("a senha deve ter entre 6 e 30 caracteres")
		}
	}

	return nil
}

func (usuario *Usuario) formatar(operacao string) error {
	usuario.Nome = strings.ToTitle(usuario.Nome)
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if operacao == "cadastro" {
		senhaComHash, erro := seguranca.HashPassword(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}
