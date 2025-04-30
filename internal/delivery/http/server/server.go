package server

import (
	"authService/internal/config"
	"authService/internal/delivery/http/middleware"
	"authService/internal/usecase"
	protos "authService/pkg/proto/gen/go"
	"context"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

func (s *Server) AuthLogin(ctx *gin.Context) {
	request := protos.GetAuthLoginRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("failed to unmarshar request: %v", err)})
		return
	}
	status, er := s.Usecase.AuthLogin(ctx, &request)
	if er != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to create course: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, &protos.GetAuthLoginResponse{Status: status})
	return
}
