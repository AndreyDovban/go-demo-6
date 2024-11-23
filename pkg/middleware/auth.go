package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := ""
		if authedHeader != "" {
			token = strings.TrimPrefix(authedHeader, "Bearer ")
		}

		fmt.Println("TOKEN", token)
		next.ServeHTTP(w, r)
	})
}
