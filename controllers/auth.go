package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"time"
)

type AuthHandler struct {
	*services.Handler
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		switch r.Method {
		case "GET":
			a.loginGet(w, r)
		case "POST":
			a.loginPost(w, r)
		}
	case "/signup":
		switch r.Method {
		case "GET":
			a.signupGet(w, r)
		case "POST":
			a.signupPost(w, r)
		}
	case "/logout":
		a.logout(w, r)
	}

}

func (a *AuthHandler) loginGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/templates/login.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		returnError(w, r, err)
	}
}

func (a *AuthHandler) loginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		returnError(w, r, err)
		return
	}
	var credentials models.LoginCredentials
	decoder := schema.NewDecoder()
	err = decoder.Decode(&credentials, r.PostForm)
	if err != nil {
		returnError(w, r, err)
		return
	}
	user, err := a.GetUserByCredentials(credentials)
	if err != nil {
		returnError(w, r, err)
		return
	}
	token, err := a.GenerateUserIDJwt(user.GetID())
	if err != nil {
		returnError(w, r, err)
		return
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func (a *AuthHandler) signupGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/templates/signup.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		returnError(w, r, err)
	}
}

func (a *AuthHandler) signupPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		returnError(w, r, err)
	}
	var user models.User
	decoder := schema.NewDecoder()
	err = decoder.Decode(&user, r.PostForm)
	if err != nil {
		returnError(w, r, err)
		return
	}
	log.Info().Msg(user.ToString())

	err = a.SaveUser(&user)
	if err != nil {
		returnError(w, r, err)
		return
	}

	token, err := a.GenerateUserIDJwt(user.GetID())
	if err != nil {
		returnError(w, r, err)
		return
	}

	log.Info().Msg(token)

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func (a *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		returnError(w, r, err)
		return
	}
	cookie.Expires = time.Now()
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}
