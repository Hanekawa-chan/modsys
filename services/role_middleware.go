package services

import (
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	Student int16 = 0
	Teacher       = 1
	Admin         = 2
)

func (h *Handler) RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/signup" || r.URL.Path == "/static" || r.URL.Path == "/error" {
			next.ServeHTTP(w, r)
			return
		}
		id, err := h.GetAuthenticatedUserID(r)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusFound)
			return
		}
		user, err := h.GetUserByID(id)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
			return
		}
		role := user.Role
		for _, path := range h.roleRoutes[role] {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}
		err = errors.New("access denied")
		log.Error().Err(err).Msg("error happened:" + err.Error())
		http.Redirect(w, r, "http://localhost:8080/error?error="+err.Error(), http.StatusFound)
	})
}
