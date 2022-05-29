package services

import (
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

const (
	Student int16 = 3
	Teacher       = 1
	Admin         = 2
)

func (h *Handler) RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		split := strings.Split(r.URL.Path, "/")
		if r.URL.Path == "/login" || r.URL.Path == "/signup" || split[1] == "static" || r.URL.Path == "/error" {
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

		for _, path := range h.roleRoutes[user.Role.Id] {
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
