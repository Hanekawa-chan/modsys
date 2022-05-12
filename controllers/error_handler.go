package controllers

import (
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
)

func returnError(w http.ResponseWriter, r *http.Request, err error) {
	log.Error().Err(err).Msg("error happened:" + err.Error())
	http.Redirect(w, r, "http://localhost:8080/error?error="+err.Error(), http.StatusFound)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/templates/error.html"))
	err := tmpl.Execute(w, map[string]string{"error": r.URL.Query().Get("error")})
	if err != nil {
		returnError(w, r, err)
	}
}
