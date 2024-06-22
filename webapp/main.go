package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/middlewares"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	r.Use(middlewares.EnableCors)

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

/*
// Meu antigo main tava assim, pelo menos essa merda rodava, mas só dava 404
func main() {
	// Serve static assets like CSS and JavaScript
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Serve HTML files
	http.Handle("/", http.FileServer(http.Dir("views")))

	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", nil)
}
*/
