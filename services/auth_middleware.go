package services

import (
	"github.com/rs/zerolog/log"
	_ "github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("path=" + r.URL.Path + " method=" + r.Method)
		if r.URL.Path == "/login" || r.URL.Path == "/signup" || r.URL.Path == "/static" || r.URL.Path == "/error" {
			next.ServeHTTP(w, r)
			return
		}
		_, err := h.GetAuthenticatedUserID(r)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
