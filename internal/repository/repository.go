package repository

import (
	"authSerivce/internal/repository/postgres"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New() fx.Option {
	return fx.Module("repository",
		fx.Provide(
			postgres.NewRepository,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, a *postgres.Repository) {
				lc.Append(fx.Hook{
					OnStart: a.OnStart,
					OnStop:  a.OnStop,
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("usecase")
		}),
	)
}
