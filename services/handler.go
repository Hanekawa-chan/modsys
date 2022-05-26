package services

import (
	"awesomeProject/dao"
	"github.com/gorilla/mux"
	"os"
)

type Handler struct {
	*mux.Router
	db         *dao.DB
	generator  *Generator
	roleRoutes map[int16][]string
}

func NewHandler(db *dao.DB, roleRoutes map[int16][]string) *Handler {
	return &Handler{
		db:         db,
		generator:  New(os.Getenv("JWT_SECRET")),
		Router:     mux.NewRouter(),
		roleRoutes: roleRoutes,
	}
}
