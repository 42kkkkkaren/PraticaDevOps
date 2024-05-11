package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta HTTP com um corpo JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		// Serialize os dados para JSON e envie como resposta HTTP
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Erro retorna uma resposta de erro HTTP com uma mensagem JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})

	// Também poderia usar algo como:
	// JSON(w, statusCode, map[string]string{"erro": mensagem})
}
