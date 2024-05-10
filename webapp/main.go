/*
package main

import (
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	utils.CarregarTemplates()
	r := router.Gerar()
	// Corrected route for serving static files
	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	// Make sure the static file handler is registered within the router
	log.Println("Serving files on http://localhost:8080/assets/")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/

package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve static files from the "webapp" directory
    fs := http.FileServer(http.Dir("./views/"))
    http.Handle("/", fs)

    log.Println("Server running on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}


