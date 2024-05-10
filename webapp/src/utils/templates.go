package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates carrega todos os templates do diretorio views
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// CarregarTemplates executa um template específico com os dados providos
func ExecutarTemplates(w http.ResponseWriter, template string, dados interface{}) error {
	erro := templates.ExecuteTemplate(w, template, dados)
	if erro != nil {
		return erro
	}

	return nil
}
