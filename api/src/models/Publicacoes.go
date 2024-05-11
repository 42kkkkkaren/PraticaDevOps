package models

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação na rede social
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadoEm,omitempty"`
}

// validar checa se a publicação é válida
func (publicacao *Publicacao) Validar() error {

	// Confere pra ver se o conteúdo está vazio
	if strings.TrimSpace(publicacao.Titulo) == "" {
		return errors.New("o título da publicação não pode estar vazio")
	}

	// Confere pra ver se o conteúdo está vazio
	if strings.TrimSpace(publicacao.Conteudo) == "" {
		return errors.New("o conteúdo da publicação não pode estar vazio")
	}

	// Verifica o tamanho do comprimento
	if len(publicacao.Conteudo) > 300 {
		return errors.New("o conteúdo da publicação não pode exceder 300 caracteres")
	}

	return nil
}
