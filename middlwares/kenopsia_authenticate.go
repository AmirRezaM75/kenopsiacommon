package middlewares

import (
	"net/http"
)

type KenopsiaAuthenticate struct {
	Token string
}

func (a KenopsiaAuthenticate) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Api-Token")

		if token == "" || token != a.Token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
