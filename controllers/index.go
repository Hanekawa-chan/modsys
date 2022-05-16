package controllers

import (
	"awesomeProject/services"
	"html/template"
	"net/http"
)

type IndexHandler struct {
	*services.Handler
}

func (i *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.indexGet(w, r)
}

func (i *IndexHandler) indexGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/templates/index.html"))
	id, err := i.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	user, err := i.GetUserByID(id)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	data := map[string]string{"name": user.Name}
	err = tmpl.Execute(w, data)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
}
