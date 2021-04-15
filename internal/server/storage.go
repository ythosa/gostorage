package server

import (
	"context"
	"fmt"
	"github.com/Ythosa/gostorage/internal/config"
	delivery2 "github.com/Ythosa/gostorage/internal/delivery"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	srv    *http.Server
	logger *logrus.Logger
}

func NewServer(router *delivery2.Handler, logger *logrus.Logger) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         fmt.Sprintf("0.0.0.0%s", config.Get().Server.BindAddr),
			WriteTimeout: time.Second * time.Duration(config.Get().Server.WriteTimeout),
			ReadTimeout:  time.Second * time.Duration(config.Get().Server.ReadTimeout),
			IdleTimeout:  time.Second * time.Duration(config.Get().Server.IdleTimeout),
			Handler:      router.R,
		},
		logger: logger,
	}
}

func (s *Server) ListenAndServe() error {
	s.logger.Printf("starting server on port: %s", config.Get().BindAddr)

	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Println("shutting down server...")

	return s.srv.Shutdown(ctx)
}
