package repository

import (
	"authService/internal/repository/postgres"
	"authService/internal/repository/redis"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New() fx.Option {
	return fx.Module("repository",
		fx.Provide(
			postgres.NewRepository,
			redis.NewRepository,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, pg *postgres.Repository, rds *redis.Repository) {
				lc.Append(fx.Hook{
					OnStart: pg.OnStart,
					OnStop:  pg.OnStop,
				})
				lc.Append(fx.Hook{
					OnStart: rds.OnStart,
					OnStop:  rds.OnStop,
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("usecase")
		}),
	)
}
