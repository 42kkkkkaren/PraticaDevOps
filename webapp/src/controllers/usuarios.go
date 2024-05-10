package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webapp/src/respostas"
)

// CriarUsuario chama a API para cadastrar um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	// Log the incoming HTTP request headers and body for debugging
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // Reset r.Body to its original state

	log.Printf("Received HTTP headers: %+v", r.Header)
	log.Printf("Received raw user data: %s", string(body))

	var userData map[string]string
	err = json.Unmarshal(body, &userData)
	if err != nil {
		log.Printf("Error unmarshalling user data: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Printf("Received user data: %s\n", usuario)

	response, erro := http.Post("http://localhost:8080/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		log.Fatal(erro)
	}
	defer response.Body.Close()

	respostas.JSON(w, response.StatusCode, nil)
}
