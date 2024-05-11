package models

// Senha representa o formato da requisição de alteração(atualização) de senha
type Senha struct {
	Nova  string `json:"senha"`
	Atual string `json:"atual"`
}
