package services

import (
	"awesomeProject/dao"
	"github.com/gorilla/mux"
	"os"
)

type Handler struct {
	*mux.Router
	db        *dao.DB
	generator *Generator
}

func NewHandler(db *dao.DB) *Handler {
	return &Handler{
		db:        db,
		generator: New(os.Getenv("JWT_SECRET")),
		Router:    mux.NewRouter(),
	}
}
