package delivery

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Err string `json:"err"`
}

func (h *Handler) error(w http.ResponseWriter, code int, err error) {
	h.respond(w, code, APIError{err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			h.respond(w, code, APIError{err.Error()})
		}
	}
}
