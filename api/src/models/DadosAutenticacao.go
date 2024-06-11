package models

// DadosAutenticacao contém o ID e o token de autenticação de um usuário
type DadosAutenticacao struct {
	ID string `json:"id"`
	Token string `json:"token"`
}