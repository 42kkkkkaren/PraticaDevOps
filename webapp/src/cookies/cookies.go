package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var secureCodec *securecookie.SecureCookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

// Configurar utiliza as variáveis de ambiente para configurar o securecookie
func Configurar() {
	secureCodec = securecookie.New(config.HashKey, config.BlockKey)
}

func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	dadosCodificados, erro := secureCodec.Encode("cookieName", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "cookieName",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Ler retorna os dados armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	// Lê o cookie com os dados do usuário
	cookie, erro := r.Cookie("cookieName")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	// Decodifica o cookie e armazena os dados em 'valores'
	if erro = secureCodec.Decode("cookieName", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}