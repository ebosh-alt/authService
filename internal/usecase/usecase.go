package usecase

import (
	"authService/internal/config"
	"authService/internal/entities"
	"authService/internal/repository/redis"
	"context"
	"time"

	protos "authService/pkg/proto/gen/go"

	"authService/internal/repository/postgres"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Usecase struct {
	cfg      *config.Config
	log      *zap.Logger
	Postgres *postgres.Repository
	Redis    *redis.Repository
	ctx      context.Context
}

func NewUsecase(
	logger *zap.Logger,
	Postgres *postgres.Repository,
	Redis *redis.Repository,
	cfg *config.Config,
	ctx context.Context,
) (*Usecase, error) {
	return &Usecase{
		cfg:      cfg,
		log:      logger,
		Postgres: Postgres,
		Redis:    Redis,
		ctx:      ctx,
	}, nil
}

func (uc *Usecase) OnStart(_ context.Context) error {
	return nil
}

func (uc *Usecase) OnStop(_ context.Context) error {
	return nil
}
func (uc *Usecase) createAccessToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"id":          user.ID,
		"telegram_id": user.TelegramID,
		"first_name":  user.FirstName,
		"exp":         time.Now().Add(365 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.cfg.Secret))
}

func (uc *Usecase) createRefreshToken(userID int32) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(10 * 365 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "signing"
	return token.SignedString([]byte(uc.cfg.Secret))
}

func (uc *Usecase) isAccessTokenValid(accessToken string) bool {
	parser := &jwt.Parser{}
	token, _, err := parser.ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		uc.log.Error("failed to parse token", zap.Error(err))
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		uc.log.Error("failed to assert token claims")
		return false
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		uc.log.Error("exp not found or invalid in token")
		return false
	}

	return time.Now().Before(time.Unix(int64(expFloat), 0))
}

func (uc *Usecase) AuthLogin(ctx context.Context, req *protos.PostAuthLoginRequest) (bool, error) {
	user := &entities.User{TelegramID: req.TelegramId, FirstName: req.FirstName}

	if err := uc.Redis.SaveCode(ctx, req.Code, user); err != nil {
		uc.log.Error("failed to save code in Redis", zap.Error(err))
		return false, err
	}

	dbUser, err := uc.Postgres.GetUserByTelegramId(ctx, user)
	if err != nil {
		uc.log.Error("failed to get user by telegram_id", zap.Error(err))
	}

	// Create user and tokens only if user doesn't exist
	if dbUser == nil {
		dbUser, err = uc.Postgres.CreateUser(ctx, user)
		if err != nil {
			uc.log.Error("failed to create user", zap.Error(err))
			return false, err
		}

		if err := uc.generateAndSaveTokens(ctx, dbUser); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (uc *Usecase) AuthVerifyCode(ctx context.Context, req *protos.PostAuthVerifyCodeRequest) (*entities.User, error) {
	userFromRedis, err := uc.Redis.VerifyCode(ctx, req.Code)
	if err != nil {
		uc.log.Error("failed to get user from Redis", zap.Error(err))
		return nil, err
	}

	dbUser, err := uc.Postgres.GetUserByTelegramId(ctx, userFromRedis)
	if err != nil {
		uc.log.Error("failed to get user from Postgres", zap.Error(err))
		return nil, err
	}

	if dbUser.TelegramID != userFromRedis.TelegramID {
		uc.log.Error("telegram_id mismatch between Redis and Postgres")
		return nil, err
	}

	return dbUser, nil
}

func (uc *Usecase) AuthRefresh(ctx context.Context, req *protos.PostAuthRefreshRequest) (*entities.User, error) {
	user, err := uc.Postgres.GetUserByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		uc.log.Error("failed to get user by refresh token", zap.Error(err))
		return nil, err
	}

	if err := uc.generateAndSaveTokens(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *Usecase) generateAndSaveTokens(ctx context.Context, user *entities.User) error {
	accessToken, err := uc.createAccessToken(user)
	if err != nil {
		uc.log.Error("failed to create access token", zap.Error(err))
		return err
	}

	refreshToken, err := uc.createRefreshToken(user.ID)
	if err != nil {
		uc.log.Error("failed to create refresh token", zap.Error(err))
		return err
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	if err := uc.Postgres.UpdateTokens(ctx, user); err != nil {
		uc.log.Error("failed to update tokens", zap.Error(err))
		return err
	}
	return nil
}
