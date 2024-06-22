package main

import (
	"api/src/config"
	"api/src/middlewares"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarEnv()
	r := router.Gerar()

	r.Use(middlewares.EnableCors)

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
