package server

import (
	"context"
	"net/http"
	"one-day-job/config"
	"one-day-job/internal/handler/Kurush"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(cfg *config.Config, handler *Kurush.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      handler.InitRoutes(),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeOut) * time.Second,
	}}
}
