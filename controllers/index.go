package controllers

import (
	"awesomeProject/services"
	"github.com/rs/zerolog/log"
	"net/http"
)

type IndexHandler struct {
	*services.Handler
}

func (i *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		switch r.Method {
		case "GET":
			i.indexGet(w, r)
		case "POST":
			i.indexPost(w, r)
		}
	}
}

func (i *IndexHandler) indexGet(w http.ResponseWriter, r *http.Request) {
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
	auth := false
	if user != nil {
		auth = true
	}

	log.Info().Msg(i.GetRole(r))

	data := map[string]interface{}{"auth": auth,
		"title": "Главная страница", "role": i.GetRole(r), "news": i.GetRecords()}

	returnTemplateWithData(w, r, "index", data)
}

func (i *IndexHandler) indexPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ReturnError(w, r, err)
		return
	}

	id := r.PostForm.Get("id")
	http.Redirect(w, r, "http://localhost:8080/test/get?id="+id, http.StatusFound)
}
