package usecase

import (
	"authService/internal/config"
	"authService/internal/entities"
	"authService/internal/repository/redis"
	"context"
	"time"

	protos "authService/pkg/proto/gen/go"

	"authService/internal/repository/postgres"
	// protos "CryptoParser/pkg/proto/gen/go"
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
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["telegram_id"] = user.TelegramID
	claims["first_name"] = user.FirstName
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, err := token.SignedString([]byte(uc.cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc *Usecase) createRefreshToken(userID int32) (string, error) {
	Claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 365 * 10).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	refreshToken.Header["kid"] = "signing"

	signedString, err := refreshToken.SignedString([]byte(uc.cfg.Secret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (uc *Usecase) AuthLogin(ctx context.Context, proto *protos.GetAuthLoginRequest) (bool, error) {
	user := &entities.User{TelegramID: proto.TelegramId, FirstName: proto.FirstName}
	er := uc.Redis.SaveCode(ctx, proto.Code, user)
	if er != nil {
		uc.log.Error("Error saving code to Redis storage:", zap.Error(er))
		return false, er
	}
	userId, er := uc.Postgres.GetUserByTelegramId(ctx, user)
	if er != nil {
		uc.log.Error("Error get user by telegram_id:", zap.Error(er))
		return false, er
	}
	if userId == 0 {
		user, er = uc.Postgres.CreateUser(ctx, user)
	} else {
		user, er = uc.Postgres.GetUserById(ctx, userId)
		if er != nil {
			uc.log.Error("Error get user by user_id:", zap.Error(er))
			return false, er
		}
	}
	accessToken, er := uc.createAccessToken(user)
	if er != nil {
		uc.log.Error("Error create access token", zap.Error(er))
		return false, er
	}
	refreshToken, er := uc.createRefreshToken(user.ID)
	if er != nil {
		uc.log.Error("Error create refresh token", zap.Error(er))
		return false, er
	}
	user.RefreshToken = refreshToken
	user.AccessToken = accessToken
	er = uc.Postgres.UpdateTokens(ctx, user)
	if er != nil {
		uc.log.Error("Error update users tokens", zap.Error(er))
		return false, er
	}
	return true, nil
}

//func (uc *Usecase) AuthVerifyCode(ctx context.Context, proto *protos.GetAuthVerifyCodeRequest) (*entities.User, error) {
//	telegramId, er := uc.Redis.VerifyCode(proto.Code)
//	if er != nil {
//		uc.log.Error("failed to verify user by code", zap.Error(er))
//	}
//
//}

//func (uc *Usecase) GetUserByJWTToken(ctx context.Context, userProto *protos.GetAuthRefreshRequest, JWTToken string) (*entities.User, error) {
//	//func (uc *Usecase) GetCourse(ctx context.Context, request *protos.GetCourseRequest, JWTToken string) (*entities.User, error) {
//	user, err := uc.Repo.GetUserByJWTToken(ctx, &entities.User{
//		RefreshToken: userProto.RefreshToken,
//	})
//	if err != nil {
//		uc.log.Error("Fail to get course info: ", zap.Error(err))
//		return nil, err
//	}
//	return user, err

//return &protos.GetCourseResponse{
//	Course: &protos.CourseCatalog{
//		CourseID:     int64(course.ID),
//		Title:        course.Title,
//		Description:  course.Description,
//		Category:     course.Category,
//		ThumbnailUrl: course.ThumbnailUrl,
//		Price:        int32(course.Price),
//		IsPaid:       &(course.IsPaid),
//	},
//	CourseContents: []*protos.CourseContent{},
//}, nil
