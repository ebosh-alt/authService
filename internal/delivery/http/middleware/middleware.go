package middleware

import (
	"authSerivce/config"
	"authSerivce/internal/repository/postgres"
	"go.uber.org/zap"
)

type Middleware struct {
	cfg   *config.Config
	repo  *postgres.Repository
	log   *zap.Logger
	roles map[string]int
}

func NewMiddleware(cfg *config.Config, log *zap.Logger, repository *postgres.Repository) *Middleware {
	return &Middleware{
		cfg:  cfg,
		log:  log,
		repo: repository,
	}
}
