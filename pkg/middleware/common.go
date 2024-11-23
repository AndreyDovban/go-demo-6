package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	SratusCode int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.SratusCode = statusCode
}
