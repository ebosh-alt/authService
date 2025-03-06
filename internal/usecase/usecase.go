package usecase

import (
	"context"
	// "fmt"
	// "math/rand"
	"authSerivce/config"

	"authSerivce/internal/repository/postgres"
	// protos "CryptoParser/pkg/proto/gen/go"
	// jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Usecase struct {
	cfg  *config.Config
	log  *zap.Logger
	Repo *postgres.Repository
	ctx  context.Context
}

const (
	clientRole    = `CLIENT`
	adminRole     = `ADMIN`
	SuccessStatus = `Success`
	BadStatus     = `Fail`
)

func NewUsecase(logger *zap.Logger, Repo *postgres.Repository, cfg *config.Config, ctx context.Context) (*Usecase, error) {
	return &Usecase{
		cfg:  cfg,
		log:  logger,
		Repo: Repo,
		ctx:  ctx,
	}, nil
}

func (uc *Usecase) OnStart(_ context.Context) error {
	//go uc.startBinanceDeamon()
	return nil
}

func (uc *Usecase) OnStop(_ context.Context) error {
	return nil
}
