package controllers

import (
	"awesomeProject/services"
	"net/http"
)

type TestHandler struct {
	*services.Handler
}

func (t *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/test":
	}
}
