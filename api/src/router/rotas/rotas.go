package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func ConfigurarRotas(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, RotaLogin)
	rotas = append(rotas, rotasSeguidores...)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	/*
		Aqui, http.FileServer(http.Dir("assets")) cria um servidor de arquivos (responsável por armazenar
		os arquivos de dados para que outros "computadores" possam acessar) pro assets
		http.Handle("/assets/", http.StripPrefix("/assets/", fs)) indica pro Go que o servidor  no diretorio
		assets quando a URL começa com /assets/. http.StripPrefix é usado pra "remover" o prefixo /assets/
		prefixo antes do servidor de arquivos olhar no diretorio.
	*/
	// Servidor de arquivos estáticos do diretorio assets (front)

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
