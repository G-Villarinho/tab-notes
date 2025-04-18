package middlewares

import (
	"net/http"

	"github.com/g-villarinho/tab-notes-api/configs"
)

func BodySizeLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, configs.Env.MaxBodySize)
		next.ServeHTTP(w, r)
	})
}
