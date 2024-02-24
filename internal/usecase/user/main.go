package user

import (
	"context"
	"errors"
	"time"

	customError "github.com/fixelti/family-hub/internal"
	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres/user"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	SignUp(ctx context.Context, email, password string) (models.UserDTO, error)
	SignIn(ctx context.Context, email, password string) (models.Tokens, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (accessToken string, err error)
}

type userUsecase struct {
	db     user.UserRepository
	config config.Config
	logger *zap.Logger
}

func New(db user.UserRepository, config config.Config, logger *zap.Logger) Usecase {
	return userUsecase{
		db:     db,
		config: config,
		logger: logger,
	}
}

func (user userUsecase) SignUp(ctx context.Context, email, password string) (models.UserDTO, error) {
	foundUser, err := user.db.GetUserByEmail(ctx, email)
	if err != nil {
		user.logger.Error("failed to get user by email", zap.Error(err))
		return models.UserDTO{}, customError.ErrInternal
	}
	if foundUser.ID != 0 {
		return models.UserDTO{}, customError.ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		user.logger.Error("failed to generate from password", zap.Error(err))
		return models.UserDTO{}, customError.ErrInternal
	}

	createdUser, err := user.db.Create(ctx, email, string(passwordHash))
	if err != nil {
		user.logger.Error("failed to create user", zap.Error(err))
		return models.UserDTO{}, customError.ErrInternal
	}

	return createdUser.ToUserDTO(), nil
}

func (user userUsecase) SignIn(ctx context.Context, email, password string) (models.Tokens, error) {
	foundUser, err := user.db.GetUserByEmail(ctx, email)
	if err != nil {
		user.logger.Error("failed to get user by email", zap.Error(err))
		return models.Tokens{}, customError.ErrInternal
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.Tokens{}, customError.ErrInvalidCredentials
		}
		user.logger.Error("failed to compare hash and password", zap.Error(err))
		return models.Tokens{}, customError.ErrInternal
	}

	// генерация access токена
	accessToken, err := generateToken(
		user.config.JWT.TokenKey,
		foundUser.ID,
		user.config.JWT.TokenLifetime,
	)

	if err != nil {
		user.logger.Error("failed to generate access token", zap.Error(err))
		return models.Tokens{}, customError.ErrInternal
	}

	// генерация refresh токена
	refreshToken, err := generateToken(
		user.config.JWT.RefreshTokenKey,
		foundUser.ID,
		user.config.JWT.RefreshTokenLifeTime,
	)

	if err != nil {
		user.logger.Error("failed to generate refresh token", zap.Error(err))
		return models.Tokens{}, customError.ErrInternal
	}

	return models.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (user userUsecase) RefreshAccessToken(ctx context.Context, refreshToken string) (accessToken string, err error) {
	refreshTokenKey := user.config.JWT.RefreshTokenKey
	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (any, error) {
		return []byte(refreshTokenKey), nil
	})
	if err != nil {
		user.logger.Error("failed to parse jwt claims", zap.Error(err))
		return accessToken, customError.ErrInternal
	}

	if !tkn.Valid {
		return accessToken, customError.ErrTokenIsNotValid
	}

	accessToken, err = generateToken(user.config.JWT.TokenKey, claims["id"].(uint), user.config.JWT.TokenLifetime)
	if err != nil {
		user.logger.Error("failed to generate access token", zap.Error(err))
		return accessToken, customError.ErrInternal
	}
	return
}

func generateToken(tokenKey string, userID uint, expirate time.Duration) (token string, err error) {
	claims := jwt.MapClaims{
		"id":       userID,
		"expirate": expirate,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString([]byte(tokenKey))
	if err != nil {
		return token, err
	}

	return token, err
}
