package controllers

import (
	"awesomeProject/models/view"
	"awesomeProject/services"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
	"net/http"
)

type AdminHandler struct {
	*services.Handler
}

func (a *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/set":
		switch r.Method {
		case "GET":
			a.setGet(w, r)
		case "POST":
			a.setPost(w, r)
		}

	}
}

func (a *AdminHandler) setGet(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"title": "Страница админа"}
	returnTemplateWithData(w, r, "admin", data)
}

func (a *AdminHandler) setPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Info().Err(err)
		ReturnError(w, r, err)
		return
	}
	var userRole view.UserRole
	decoder := schema.NewDecoder()
	err = decoder.Decode(&userRole, r.PostForm)
	if err != nil {
		log.Info().Err(err)
		ReturnError(w, r, err)
		return
	}

	err = a.SetRole(userRole.Email, userRole.Role)
	if err != nil {
		ReturnError(w, r, err)
	}
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}
