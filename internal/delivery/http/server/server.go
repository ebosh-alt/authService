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
	log        *zap.Logger
	cfg        *config.Config
	serv       *gin.Engine
	Usecase    *usecase.Usecase
	middleware *middleware.Middleware
}

func NewServer(logger *zap.Logger, cfg *config.Config, uc *usecase.Usecase, middleware *middleware.Middleware) (*Server, error) {
	return &Server{
		log:        logger,
		cfg:        cfg,
		serv:       gin.Default(),
		Usecase:    uc,
		middleware: middleware,
	}, nil
}

func (s *Server) OnStart(_ context.Context) error {
	s.createController()
	go func() {
		s.log.Debug("serv started")
		if err := s.serv.Run(s.cfg.Server.Host + ":" + s.cfg.Server.Port); err != nil {
			s.log.Error("failed to serve: " + err.Error())
		}
		return
	}()
	return nil
}

func (s *Server) OnStop(_ context.Context) error {
	s.log.Debug("stop gRPS")
	//s.serv.GracefulStop()
	return nil
}

func (s *Server) AuthLogin(ctx *gin.Context) {
	var req protos.PostAuthLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error("failed to bind AuthLogin request", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("invalid request payload: %v", err),
		})
		return
	}

	status, err := s.Usecase.AuthLogin(ctx, &req)
	if err != nil {
		s.log.Error("usecase AuthLogin failed", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, &protos.PostAuthLoginResponse{Status: status})
}

func (s *Server) AuthVerifyCode(ctx *gin.Context) {
	var req protos.PostAuthVerifyCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error("failed to bind AuthLogin request", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("invalid request payload: %v", err),
		})
		return
	}
	user, err := s.Usecase.AuthVerifyCode(ctx, &req)
	if err != nil {
		s.log.Error("usecase AuthVerifyCode failed", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return

	}
	ctx.JSON(http.StatusOK, &protos.PostAuthVerifyCodeResponse{
		Status:       true,
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	})
}

func (s *Server) AuthRefresh(ctx *gin.Context) {
	var req protos.PostAuthRefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error("failed to bind AuthRefresh request", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("invalid request payload: %v", err),
		})
		return
	}
	user, err := s.Usecase.AuthRefresh(ctx, &req)
	if err != nil {
		s.log.Error("usecase AuthRefresh failed", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, &protos.PostAuthRefreshResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	})
}
