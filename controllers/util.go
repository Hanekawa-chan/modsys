package controllers

import (
	"html/template"
	"net/http"
)

func returnTemplate(w http.ResponseWriter, r *http.Request, templateName string) {
	tmpl := template.Must(template.ParseFiles("view/templates/" + templateName + ".html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		ReturnError(w, r, err)
	}
}
