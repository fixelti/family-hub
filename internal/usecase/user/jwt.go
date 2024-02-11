package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (user userUsecase) SignUp(ctx context.Context, email, password string) (uint, error) {
	userID, err := user.db.GetIDByEmail(ctx, email)
	if err != nil {
		return 0, err
	}
	if userID != 0 {
		return 0, ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return 0, err
	}

	id, err := user.db.Create(ctx, email, string(passwordHash))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (user userUsecase) SignIn(ctx context.Context, email, password string) (models.Tokens, error) {
	usr, err := user.db.GetIDAndPasswordrByEmail(ctx, email)
	if err != nil {
		log.Printf("failed to get user by email: %s", err)
		return models.Tokens{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)); err != nil {
		return models.Tokens{}, ErrInvalidCredentials
	}

	accessToken, err := generateAccessToken(
		user.config.Environment.TokenKey,
		models.Payload{UserID: usr.ID, Expirate: user.config.JWT.TokenLifetime},
	)

	if err != nil {
		log.Printf("failed to generate access token: %s", err)
		return models.Tokens{}, err
	}

	refreshToken, err := generateRefreshToken(
		user.config.Environment.RefreshTokenKey,
		models.Payload{Expirate: user.config.JWT.RefreshTokenLifeTime},
	)

	if err != nil {
		log.Printf("failed to generate refresh token: %s", err)
		return models.Tokens{}, err
	}

	return models.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (user userUsecase) RefreshAccessToken(ctx context.Context, refreshToken string) (accessToken string, err error) {
	refreshTokenKey := user.config.Environment.RefreshTokenKey
	claims := models.Payload{}
	fmt.Println(refreshTokenKey)
	tkn, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (any, error) {
		return []byte(refreshTokenKey), nil
	})
	if err != nil {
		log.Printf("failed to parse with claims: %s\n", err)
		return accessToken, err
	}

	if !tkn.Valid {
		log.Println("failed to refresh token: token is not valid")
		return accessToken, ErrTokenIsNotValid
	}

	claims.Expirate = user.config.JWT.TokenLifetime
	accessToken, err = generateAccessToken(user.config.Environment.TokenKey, claims)
	return
}

func generateAccessToken(tokenKey string, payload models.Payload) (string, error) {
	jwtPayload := jwt.MapClaims{
		"user_id":  payload.UserID,
		"expirate": time.Now().Add(time.Minute * time.Duration(payload.Expirate)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	strToken, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		log.Printf("failed to signed string token: %s", err)
		return "", err
	}

	return strToken, nil
}

func generateRefreshToken(tokenKey string, payload models.Payload) (string, error) {
	jwtPayload := jwt.MapClaims{
		"expirate": time.Now().Add(time.Minute * time.Duration(payload.Expirate)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	strToken, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		log.Printf("failed to signed string token: %s", err)
		return "", err
	}

	return strToken, nil
}
