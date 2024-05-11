package seguranca

import "golang.org/x/crypto/bcrypt"

// HashPassword recebe uma string e coloca um hash nela
func HashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes), err
}

// CheckPasswordHash compara uma senha e um hash e retorna se elas são iguais
func CheckPasswordHash(hash string, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}
