package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco conexão com o MySQL
	StringConexaoBanco = ""

	// Porta aonde a API vai estar rodando
	Porta = 0

	// SecretKey é a chave que será usada para assinar o Token
	SecretKey []byte
)

func CarregarEnv() {
	var erro error
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	if string(SecretKey) == "" {
		log.Fatal("JWT_SECRET is not set in .env file")
	}

}
