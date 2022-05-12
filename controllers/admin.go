package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"github.com/gorilla/schema"
	"net/http"
)

type AdminHandler struct {
	*services.Handler
}

func (a *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/set":
		a.setRole(w, r)
	}
}

func (a *AdminHandler) setRole(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		returnError(w, r, err)
		return
	}
	var userRole models.UserRole
	decoder := schema.NewDecoder()
	err = decoder.Decode(&userRole, r.PostForm)
	if err != nil {
		returnError(w, r, err)
		return
	}

	err = a.SetRole(userRole.Email, userRole.Role)
	if err != nil {
		returnError(w, r, err)
		return
	}
}
