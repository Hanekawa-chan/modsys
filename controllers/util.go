package controllers

import (
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
)

func returnTemplate(w http.ResponseWriter, r *http.Request, templateName string) {
	tmpl := template.Must(template.ParseFiles("view/templates/" + templateName + ".html"))
	_, err := tmpl.ParseFiles("view/templates/base.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		ReturnError(w, r, err)
	}
}

func returnTemplateWithData(w http.ResponseWriter, r *http.Request, templateName string, data any) {
	tmpl, err := template.ParseFiles("view/templates/base.html", "view/templates/"+templateName+".html")
	if err != nil {
		ReturnError(w, r, err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ReturnError(w, r, err)
	}
}

func ReturnError(w http.ResponseWriter, r *http.Request, err error) {
	log.Error().Err(err).Msg("error happened:" + err.Error())
	http.Redirect(w, r, "http://localhost:8080/error?error="+err.Error(), http.StatusFound)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"error": r.URL.Query().Get("error"), "title": "Ошибка"}
	returnTemplateWithData(w, r, "error", data)
}
