package postgres

import (
	"authSerivce/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	ctx context.Context
	log *zap.Logger
	cfg *config.Config
	DB  *pgxpool.Pool
}

func NewRepository(log *zap.Logger, cfg *config.Config, ctx context.Context) (*Repository, error) {
	return &Repository{
		ctx: ctx,
		log: log,
		cfg: cfg,
	}, nil
}

func (r *Repository) OnStart(_ context.Context) error {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		r.cfg.Postgres.Host,
		r.cfg.Postgres.Port,
		r.cfg.Postgres.User,
		r.cfg.Postgres.Password,
		r.cfg.Postgres.DBName,
		r.cfg.Postgres.SSLMode)

	r.log.Info(connectionUrl)
	pool, err := pgxpool.Connect(r.ctx, connectionUrl)
	if err != nil {
		return err
	}
	r.DB = pool
	return nil
}

func (r *Repository) OnStop(_ context.Context) error {
	r.DB.Close()
	return nil
}
