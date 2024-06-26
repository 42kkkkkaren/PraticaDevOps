package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login handler called") // Log para verificar se a função está sendo chamada
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Println("Erro ao ler corpo da requisição:", erro)
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		log.Println("Erro ao fazer unmarshal do JSON:", erro)
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarDB()
	if erro != nil {
		log.Println("Erro ao conectar no banco de dados:", erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	log.Println("Buscando usuário no banco de dados:", usuario.Email)
	// Busca pelo usuário usando o email fornecido
	repositorio := repositorios.NovoUsuarioRepositorio(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		log.Println("Erro ao buscar usuário no banco de dados:", erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	log.Println("Usuário encontrado:", usuarioSalvoNoBanco.Email)

	// Verifica se a senha fornecida corresponde à senha hash armazenada no banco de dados
	if erro = seguranca.CheckPasswordHash(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		// Se as senhas não corresponderem, retorna erro 401 (não autorizado)
		log.Println("Senha incorreta:", erro)
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	log.Println("Senha verificada com sucesso")

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		log.Println("Erro ao criar token:", erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	log.Println("Token criado com sucesso")

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)
	respostas.JSON(w, http.StatusOK, models.DadosAutenticacao{ID: usuarioID, Token: token})

}
