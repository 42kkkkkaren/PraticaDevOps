package router

import (
	"api/src/router/rotas"
	"net/http"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um Router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	// Adiciona uma rota OPTIONS para todas as rotas
	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})

	return rotas.ConfigurarRotas(r)
}
