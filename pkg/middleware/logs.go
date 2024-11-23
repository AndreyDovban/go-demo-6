package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			SratusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)
		fmt.Println(wrapper.SratusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
