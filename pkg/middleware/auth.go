package middleware

import (
	"go-demo-6/configs"
	"go-demo-6/pkg/jwt"
	"net/http"
	"strings"

	"golang.org/x/net/context"
)

type key string

const ContextEmailKey key = "ContextEmailKey"

func wrinteUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			wrinteUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			wrinteUnauthed(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
