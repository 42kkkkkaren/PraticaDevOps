package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do usuário do contexto
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// Extrait o ID do usuário alvo para seguir, pelos parametros URL
	parametros := mux.Vars(r)
	seguindoID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Verificar se o usuário não está tentando seguir ele mesmo :(
	if seguidorID == seguindoID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir o prórpio usuário"))
		return
	}

	// Após as etapas acima, podemos abrir a conexão com o Banco de Dados
	db, erro := banco.ConectarDB()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoSeguidorRepositorio(db)
	if erro = repositorio.Seguir(seguidorID, seguindoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUsuario permite que um usuário pare de seguir outro
func UnfollowUsuario(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do usuário do contexto
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// Extrair o ID do usuário alvo para deixar de seguir, pelos parametros URL
	parametros := mux.Vars(r)
	seguindoID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Verificar se o usuário não está tentando deixar seguir ele mesmo :(
	if seguidorID == seguindoID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deixar de seguir o prórpio usuário"))
		return
	}

	// Após as etapas acima, podemos abrir a conexão com o Banco de Dados
	db, erro := banco.ConectarDB()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoSeguidorRepositorio(db)
	if erro = repositorio.Unfollow(seguidorID, seguindoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores traz todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do usuário alvo para ver quem segue ele, pelos parametros URL
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarDB()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoSeguidorRepositorio(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)

}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do usuário alvo para ver quem ele segue, pelos parametros URL
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarDB()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoSeguidorRepositorio(db)
	seguindo, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguindo)
}
