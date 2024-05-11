package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// UsuarioRepository representa o repositório para a entidade Usuario
type Usuarios struct {
	db *sql.DB
}

// NovoUsuarioRepositorio cria uma nova instância de UsuarioRepository
func NovoUsuarioRepositorio(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um Usuario no Banco de Dados
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar retorna os usuarios que atendem a um filtro de Nome ou Nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	// Adiciona os caracteres de porcentagem para busca parcial
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	// Prepara a query SQL para buscar usuários que correspondam ao nome ou nick fornecido
	query := "SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?"

	linhas, erro := repositorio.db.Query(query, nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	// Itera sobre os resultados da query
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID retorna o usuario do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	query := "SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?"

	linhas, erro := repositorio.db.Query(query, ID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	// Um slice de usuarios para segurar os dados retornados das linhas
	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar modifica um usuário existente no banco de dados com base no ID fornecido
func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {

	query := "UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?"
	statement, erro := repositorio.db.Prepare(query)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)
	if erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {

	// Verifica se o usuário existe
	if _, err := repositorio.BuscarPorID(ID); err != nil {
		return err // Retorna o erro se o usuário não existir
	}

	query := "DELETE FROM usuarios WHERE id = ?"
	_, erro := repositorio.db.Exec(query, ID)
	if erro != nil {
		return erro
	}
	return nil
}

// BuscarPorEmail busca um usuário por emaail e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	query := "select id, senha from usuarios where email = ?"
	linha, erro := repositorio.db.Query(query, email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// BuscarSenha traz a senha de um usuario pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {

	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil

	/* Sepa dava pra fazer desse jeito aq, tenho q averiguar melhor
	var senha string
	erro := repositorio.db.QueryRow("SELECT senha FROM usuarios WHERE id = ?", usuarioID).Scan(&senha)
	return senha, erro
	*/
}

// AtualizarSenha alter a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}
