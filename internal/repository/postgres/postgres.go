package postgres

import (
	"authService/internal/config"
	"authService/internal/entities"
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

const getUserByTelegramId = `
	SELECT id, telegram_id, first_name, access_token, refresh_token, created_at, updated_at FROM users WHERE telegram_id = $1
`

func (r *Repository) GetUserByTelegramId(ctx context.Context, user *entities.User) (*entities.User, error) {
	userDTO := user.ConvertToDTO()
	err := r.DB.QueryRow(ctx, getUserById, user.TelegramID).Scan(
		&userDTO.ID,
		&userDTO.TelegramID,
		&userDTO.FirstName,
		&userDTO.AccessToken,
		&userDTO.RefreshToken,
		&userDTO.CreateAt,
		&userDTO.UpdatedAt,
	)
	if err != nil {
		r.log.Error("fail to get user by telegram id", zap.Error(err))
		return nil, err
	}
	return userDTO.FromDTOConvert(), nil
}

const getUserById = `
	SELECT id, telegram_id, first_name, access_token, refresh_token, created_at, updated_at
	               FROM users WHERE id = $1
`

func (r *Repository) GetUserById(ctx context.Context, id int32) (*entities.User, error) {
	var userDTO entities.UserDTO
	err := r.DB.QueryRow(ctx, getUserById, id).Scan(
		&userDTO.ID,
		&userDTO.TelegramID,
		&userDTO.FirstName,
		&userDTO.AccessToken,
		&userDTO.RefreshToken,
		&userDTO.CreateAt,
		&userDTO.UpdatedAt,
	)
	if err != nil {
		r.log.Error("fail to get user by id", zap.Error(err))
		return nil, err
	}
	return userDTO.FromDTOConvert(), nil
}

const getUserByRefreshToken = `
	SELECT id, telegram_id, first_name, access_token, refresh_token, created_at, updated_at FROM users WHERE refresh_token = $1
`

func (r *Repository) GetUserByRefreshToken(ctx context.Context, refreshToken string) (*entities.User, error) {
	var userDTO entities.UserDTO
	err := r.DB.QueryRow(ctx, getUserById, refreshToken).Scan(
		&userDTO.ID,
		&userDTO.TelegramID,
		&userDTO.FirstName,
		&userDTO.AccessToken,
		&userDTO.RefreshToken,
		&userDTO.CreateAt,
		&userDTO.UpdatedAt,
	)
	if err != nil {
		r.log.Error("fail to get user_id by refresh_token", zap.Error(err))
		return nil, err
	}
	return userDTO.FromDTOConvert(), nil
}

const createUser = `
	INSERT INTO users (telegram_id, first_name) values ($1, $2) returning id
`

func (r *Repository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	userDTO := user.ConvertToDTO()

	er := r.DB.QueryRow(ctx, createUser, userDTO.TelegramID, userDTO.FirstName).Scan(
		&userDTO.ID,
	)
	if er != nil {
		r.log.Error("fail to create user", zap.Error(er))
		return nil, er
	}
	return userDTO.FromDTOConvert(), nil

}

const updateTokens = `
	UPDATE users SET access_token = $1, refresh_token = $2 WHERE id = $3
`

func (r *Repository) UpdateTokens(ctx context.Context, user *entities.User) error {
	userDTO := user.ConvertToDTO()
	_, er := r.DB.Exec(ctx, updateTokens, userDTO.AccessToken, userDTO.RefreshToken, userDTO.ID)
	if er != nil {
		r.log.Error("failed update tokens user", zap.Error(er))
		return er
	}
	return nil
}
