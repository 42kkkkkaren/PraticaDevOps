package respostas

import (
	"encoding/json"
	"net/http"
)

// JSON retorna uma resposta em formato JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dados.([]byte))

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		http.Error(w, "Houve um erro interno no servidor", http.StatusInternalServerError)
		return
	}
}