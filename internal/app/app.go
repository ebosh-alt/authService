package app

import (
	"authSerivce/config"
	"authSerivce/internal/delivery/http/middleware"
	"authSerivce/internal/delivery/http/server"
	"authSerivce/internal/repository"
	"authSerivce/internal/usecase"
	"context"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func New() *fx.App {
	return fx.New(
		fx.Options(
			repository.New(),
			usecase.New(),
			middleware.New(),
			server.New(),
		),
		fx.Provide(
			context.Background,
			config.NewConfig,
			zap.NewDevelopment,
		),
		fx.WithLogger(
			func(log *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: log}
			},
		),
	)
}
