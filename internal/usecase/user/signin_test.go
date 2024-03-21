package user_test

import (
	"context"
	"testing"
	"time"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/usecase"
	jwtLib "github.com/fixelti/family-hub/lib/jwt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestSignIn(t *testing.T) {
	t.Parallel()
	t.Helper()

	const (
		email    = "test@gmail.com"
		password = "testPassword"

		accessTokenKey                     = "testAccess"
		refreshTokenKey                    = "testRefresh"
		expirateAccessToken  time.Duration = 1 * time.Hour
		expirateRefreshToken time.Duration = 5 * time.Hour
	)

	jwtLib := jwtLib.New(accessTokenKey, refreshTokenKey, expirateAccessToken, expirateRefreshToken)

	t.Run("Success", func(t *testing.T) {
		mocks := NewSuite(t)

		var (
			expectedUserID                    = 1
			getUserByEmailResponse            = models.User{ID: 1, Password: password}
			getUserByEmailResponseError error = nil

			compareHashAndPasswordResponseError error = nil
		)

		mocks.UserRepository.
			EXPECT().
			GetUserByEmail(gomock.Any(), email).
			Return(getUserByEmailResponse, getUserByEmailResponseError)

		mocks.Crypto.
			EXPECT().
			CompareHashAndPassword([]byte(getUserByEmailResponse.Password), []byte(password)).
			Return(compareHashAndPasswordResponseError)

		usecase := usecase.New(config.Config{}, nil, mocks.UserRepository, nil, mocks.Crypto, jwtLib)
		tokens, err := usecase.User.SignIn(context.Background(), email, password)
		assert.NoError(t, err)

		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(tokens.Access, &claims, func(token *jwt.Token) (any, error) {
			return []byte(accessTokenKey), nil
		})
		assert.NoError(t, err)

		assert.Equal(t, expectedUserID, int(claims["id"].(float64)))

		_, err = jwt.ParseWithClaims(tokens.Refresh, &claims, func(token *jwt.Token) (any, error) {
			return []byte(refreshTokenKey), nil
		})
		assert.NoError(t, err)

		assert.Equal(t, expectedUserID, int(claims["id"].(float64)))
	})

	t.Run("Wrong login", func(t *testing.T) {
		mocks := NewSuite(t)

		var (
			expectedError               error = customError.ErrInvalidCredentials
			getUserByEmailResponse            = models.User{ID: 0, Password: ""}
			getUserByEmailResponseError error = nil

			wrongEmail = "test2@gmail.com"
		)

		mocks.UserRepository.
			EXPECT().
			GetUserByEmail(gomock.Any(), wrongEmail).
			Return(getUserByEmailResponse, getUserByEmailResponseError)

		usecase := usecase.New(config.Config{}, nil, mocks.UserRepository, nil, mocks.Crypto, jwtLib)
		_, err := usecase.User.SignIn(context.Background(), wrongEmail, password)
		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("Wrong password", func(t *testing.T) {
		mocks := NewSuite(t)

		var (
			expectedError               error = customError.ErrInvalidCredentials
			getUserByEmailResponse            = models.User{ID: 1, Password: password}
			getUserByEmailResponseError error = nil

			compareHashAndPasswordResponseError error = bcrypt.ErrMismatchedHashAndPassword

			wrongPassword = "wrongPassword"
		)

		mocks.UserRepository.
			EXPECT().
			GetUserByEmail(gomock.Any(), email).
			Return(getUserByEmailResponse, getUserByEmailResponseError)

		mocks.Crypto.
			EXPECT().
			CompareHashAndPassword([]byte(getUserByEmailResponse.Password), []byte(wrongPassword)).
			Return(compareHashAndPasswordResponseError)

		usecase := usecase.New(config.Config{}, nil, mocks.UserRepository, nil, mocks.Crypto, jwtLib)
		_, err := usecase.User.SignIn(context.Background(), email, wrongPassword)
		assert.ErrorIs(t, err, expectedError)
	})
}
