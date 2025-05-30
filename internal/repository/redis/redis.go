package redis

import (
	"authService/internal/config"
	"authService/internal/entities"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type Repository struct {
	ctx    context.Context
	log    *zap.Logger
	cfg    *config.Config
	Client *redis.Client
}

func NewRepository(log *zap.Logger, cfg *config.Config, ctx context.Context) (*Repository, error) {
	return &Repository{
		ctx: ctx,
		log: log,
		cfg: cfg,
	}, nil
}

func (r *Repository) OnStart(_ context.Context) error {
	addr := fmt.Sprintf("%s:%s", r.cfg.Redis.Host, r.cfg.Redis.Port)
	r.log.Info("connecting to Redis", zap.String("address", addr))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.cfg.Redis.Password,
		DB:       r.cfg.Redis.DB,
	})

	if err := client.Ping(r.ctx).Err(); err != nil {
		return err
	}

	r.Client = client
	return nil
}

func (r *Repository) OnStop(_ context.Context) error {
	if err := r.Client.Close(); err != nil {
		r.log.Error("error closing Redis connection", zap.Error(err))
		return err
	}
	return nil
}

func (r *Repository) SaveCode(ctx context.Context, code string, user *entities.User) error {
	key := fmt.Sprintf("auth_code:%s", code)

	userData, err := json.Marshal(user)
	if err != nil {
		r.log.Error("failed to marshal user data", zap.Error(err))
		return err
	}

	err = r.Client.Set(ctx, key, userData, 60*time.Minute).Err()
	if err != nil {
		r.log.Error("failed to save code in Redis", zap.Error(err))
		return err
	}

	r.log.Info("saved code in Redis", zap.String("code", code), zap.Int64("telegramID", user.TelegramID))
	return nil
}

func (r *Repository) VerifyCode(ctx context.Context, code string) (*entities.User, error) {
	key := fmt.Sprintf("auth_code:%s", code)

	data, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			r.log.Warn("code not found in Redis", zap.String("code", code))
			return nil, nil
		}
		r.log.Error("failed to get code from Redis", zap.Error(err))
		return nil, err
	}

	var user entities.User
	if err := json.Unmarshal(data, &user); err != nil {
		r.log.Error("failed to unmarshal user data", zap.Error(err))
		return nil, err
	}

	return &user, nil
}
