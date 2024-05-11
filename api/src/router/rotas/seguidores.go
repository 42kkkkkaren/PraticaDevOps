package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasSeguidores = []Rota{
	{
		URI:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UnfollowUsuario,
		RequerAutenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
}
