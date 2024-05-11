package repositorios

import (
	"api/src/models"
	"database/sql"
)

// Publicacoes representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovaPublicacaoRepositorio cria um repositorio de publicações
func NovaPublicacaoRepositorio(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma unica publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {

	/*
			p.* seleciona todas as colunas da tabela de publicacoes
		    u.nick seleciona o a coluna nick da tabela de usuarios
		    O INNER JOIN obtém linhas que tem valores correspondentes em ambas as tabelas
		    WHERE p.id = ? filtra o resultado para incluir apenas a publicação com o ID especifico
	*/
	linha, erro := repositorio.db.Query(`
		SELECT p.*, u.nick FROM publicacoes p inner join usuarios u on u.id = p.autor_id WHERE p.id = ?`,
		publicacaoID)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao
	if linha.Next() {
		erro := linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		)
		if erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publicacoes p
		JOIN usuarios u ON p.autor_id = u.id
		LEFT JOIN seguidores s ON u.id = s.seguindo_id AND s.seguidor_id = ?
		WHERE p.autor_id = ? OR s.seguidor_id = ?
		ORDER BY p.criadaEm DESC`,
		usuarioID, usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	// Itera sobre o conjunto de resultados e preenche o slice de publicacoes
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Atualizar altera os dados de uma publicação no banco
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {

	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario traz todas as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, u.nick from publicacoes p
		JOIN usuarios u on u.id = p.autor_id
		WHERE p.autor_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// IncrementarCurtida adiciona uma curtida na publicação
func (repositorio Publicacoes) IncrementarCurtida(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) DecrementarCurtida(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE publicacoes SET curtidas = 
		CASE 
			WHEN curtidas > 0 THEN curtidas - 1 
			ELSE 0 
		END 
		WHERE id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}
