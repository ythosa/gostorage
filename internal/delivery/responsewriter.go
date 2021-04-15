package delivery

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	Code int
}

// WriteHeader ...
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
