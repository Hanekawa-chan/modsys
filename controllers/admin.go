package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	*services.Handler
}

func (a *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/set":
		a.setGet(w, r)
	case "/list":
		a.listGet(w, r)
	case "/record":
		switch r.Method {
		case "GET":
			a.recordGet(w, r)
		case "POST":
			a.recordPost(w, r)
		}
	}
}

func (a *AdminHandler) setGet(w http.ResponseWriter, r *http.Request) {
	qId := r.URL.Query().Get("id")
	qRole := r.URL.Query().Get("role")
	userId, err := uuid.Parse(qId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	user, err := a.GetUserByID(userId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	role, err := strconv.Atoi(qRole)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	err = a.SetRole(user.Email, int16(role))
	if err != nil {
		ReturnError(w, r, err)
		return
	}
}

func (a *AdminHandler) listGet(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{"title": "Список пользователей", "auth": true,
		"users": a.GetUsers(), "role": a.GetRole(r)}
	returnTemplateWithData(w, r, "list", data)
}

func (a *AdminHandler) recordGet(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{"title": "Создание статьи", "auth": true, "role": a.GetRole(r)}
	returnTemplateWithData(w, r, "record", data)
}

func (a *AdminHandler) recordPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Info().Err(err)
		ReturnError(w, r, err)
		return
	}
	var record models.Record
	decoder := schema.NewDecoder()
	err = decoder.Decode(&record, r.PostForm)
	if err != nil {
		log.Info().Err(err)
		ReturnError(w, r, err)
		return
	}

	err = a.SaveRecord(record)
	if err != nil {
		ReturnError(w, r, err)
	}
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}
