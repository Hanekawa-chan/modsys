package services

import (
	"errors"
	"github.com/rs/zerolog/log"
	_ "github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
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
		log.Info().Msg("Authenticated user. ID = " + id.String())
		if r.URL.Path == "/set" {
			user, err := h.GetUserByID(id)
			if err != nil {
				http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
				return
			}
			if user.Role != 2 {
				err = errors.New("access denied")
				log.Error().Err(err).Msg("error happened:" + err.Error())
				http.Redirect(w, r, "http://localhost:8080/error?error="+err.Error(), http.StatusFound)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
