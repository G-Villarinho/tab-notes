package middlewares

import (
	"net/http"
	"strings"

	"github.com/hermes-mailer/config"
)

func RequireAPIKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-API-Key")
			if key == "" {
				authHeader := r.Header.Get("Authorization")
				if strings.HasPrefix(authHeader, "Bearer ") {
					key = strings.TrimPrefix(authHeader, "Bearer ")
				}
			}

			if key != config.Env.API.APIKey {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
