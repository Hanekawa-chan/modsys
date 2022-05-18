package controllers

import (
	"awesomeProject/services"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

type TestHandler struct {
	*services.Handler
}

func (t *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/test/create":
		switch r.Method {
		case "GET":
			t.testCreateGet(w, r)
		case "POST":
			t.testCreatePost(w, r)
		}
	case "/test/get":
		switch r.Method {
		case "GET":
			t.testGet(w, r)
		case "POST":
			t.testPost(w, r)
		}
	}
}

func (t *TestHandler) testCreateGet(w http.ResponseWriter, r *http.Request) {
	returnTemplate(w, r, "test_create")
}

func (t *TestHandler) testCreatePost(w http.ResponseWriter, r *http.Request) {

}

func (t *TestHandler) testGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/templates/test.html"))
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	test, err := t.GetTestByID(id)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	data := map[string]string{"name": test.Name}
	err = tmpl.Execute(w, data)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
}

func (t *TestHandler) testPost(w http.ResponseWriter, r *http.Request) {

}
