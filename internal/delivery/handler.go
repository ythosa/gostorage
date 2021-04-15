package delivery

import (
	"github.com/Ythosa/gostorage/internal/service"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Ythosa/gostorage/docs" // Generated swagger documentation
)

type Handler struct {
	R       *mux.Router
	logger  *logrus.Logger
	service *service.Service
}

func NewHandler(logger *logrus.Logger, r *mux.Router, service *service.Service) *Handler {
	return &Handler{
		R:       r,
		logger:  logger,
		service: service,
	}
}

func (h *Handler) ConfigureRouter() {
	h.R.Use(h.setRequestID)
	h.R.Use(h.logRequest)
	h.R.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// Swagger documentation
	h.R.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// File Server
	static := h.R.PathPrefix("/static/").Subrouter()
	static.PathPrefix("").Handler(http.FileServer(http.Dir("")))

	api := h.R.PathPrefix("/api").Subrouter()

	file := api.PathPrefix("/file").Subrouter()
	file.HandleFunc("/upload", h.uploadFile()).Methods("POST")
	file.HandleFunc("/update", h.updateFile()).Methods("PATCH")
	file.HandleFunc("/delete", h.deleteFile()).Methods("DELETE")
}
