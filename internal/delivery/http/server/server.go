package server

import (
	"authSerivce/config"
	"authSerivce/internal/delivery/http/middleware"
	"authSerivce/internal/usecase"
	"context"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "net/http"
)

type Server struct {
	logger     *zap.Logger
	cfg        *config.Config
	serv       *gin.Engine
	Usecase    *usecase.Usecase
	middleware *middleware.Middleware
}

func NewServer(logger *zap.Logger, cfg *config.Config, uc *usecase.Usecase, middleware *middleware.Middleware) (*Server, error) {
	return &Server{
		logger:     logger,
		cfg:        cfg,
		serv:       gin.Default(),
		Usecase:    uc,
		middleware: middleware,
	}, nil
}

func (s *Server) OnStart(_ context.Context) error {
	s.createController()
	go func() {
		s.logger.Debug("serv started")
		if err := s.serv.Run(s.cfg.Server.Host + ":" + s.cfg.Server.Port); err != nil {
			s.logger.Error("failed to serve: " + err.Error())
		}
		return
	}()
	return nil
}

func (s *Server) OnStop(_ context.Context) error {
	s.logger.Debug("stop gRPS")
	//s.serv.GracefulStop()
	return nil
}
