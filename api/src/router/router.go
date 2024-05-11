package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um Router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.ConfigurarRotas(r)
}
