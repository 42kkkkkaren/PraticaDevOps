package repositorios

import (
	"api/src/models"
	"database/sql"
)

// Seguidores struct representa o repositorio para a funcionalidade dos seguidores
type Seguidores struct {
	db *sql.DB
}

// NovoUsuarioRepositorio cria uma nova instância de UsuarioRepository
func NovoSeguidorRepositorio(db *sql.DB) *Seguidores {
	return &Seguidores{db}
}

// Seguir permite que um usuário siga outro
func (repositorio Seguidores) Seguir(seguidorID, seguindoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (seguidor_id, seguindo_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(seguidorID, seguindoID)
	if erro != nil {
		return erro
	}
	return nil
}

func (repositorio Seguidores) Unfollow(seguidorID, seguindoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where seguidor_id = ? and seguindo_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(seguidorID, seguindoID)
	if erro != nil {
		return erro
	}
	return nil
}

/*
Nesse código, BuscarSeguidores é um método da estrutura Seguidores que executa uma consulta SQL
 para buscar todos os usuários que seguem um determinado usuário (usuarioID). Ele retorna uma fatia de structs
 Usuario representando os seguidores, juntamente com um erro se a operação falhar em algum ponto. O uso do deferimento
 garante que os recursos sejam limpos, evitando possíveis vazamentos de memória.
*/
// BuscarSeguidores executa uma query para achar todos os seguidores de um dado usuário
func (repositorio Seguidores) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
	FROM usuarios u INNER JOIN seguidores s ON s.seguidor_id = u.id 
	WHERE s.seguindo_id = ?
	`, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []models.Usuario // Slice para manter os seguidores

	// Iterar sobre o resultado da query e popular o slice de seguidores
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, usuario)
	}
	return seguidores, nil // Retorna o slice populado e nil error se der certo
}

// BuscarSeguindo executa uma query para achar todos os usuarios que um usuário está seguindo
func (repositorio Seguidores) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
	FROM usuarios u INNER JOIN seguidores s ON s.seguindo_id = u.id 
	WHERE s.seguidor_id = ?
	`, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguindo []models.Usuario // Slice para manter os seguindo

	// Iterar sobre o resultado da query e popular o slice de seguindo
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		seguindo = append(seguindo, usuario)
	}
	return seguindo, nil // Retorna o slice populado e nil error se der certo
}
