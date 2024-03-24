package user_test

import (
	"context"
	"testing"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSignUp(t *testing.T) {
	t.Parallel()
	t.Helper()

	const (
		email    = "test@gmail.com"
		password = "testPassword"
	)

	t.Run("Success", func(t *testing.T) {
		mocks := NewSuite(t)

		var (
			expectedUser              = models.UserDTO{ID: 1, Email: email}
			getUserByEmailUser        = models.User{ID: 0, Email: ""}
			createUser                = models.User{ID: 1, Email: email, Password: ""}
			createError         error = nil
			getUserByEmailError error = nil
		)

		mocks.UserRepository.
			EXPECT().
			GetUserByEmail(gomock.Any(), email).
			Return(getUserByEmailUser, getUserByEmailError)

		mocks.UserRepository.
			EXPECT().
			Create(gomock.Any(), email, password).
			Return(createUser, createError)

		mocks.Crypto.
			EXPECT().
			GenerateFromPassword([]byte(password), gomock.Any()).
			Return([]byte(password), nil)

		usecase := usecase.New(config.Config{}, nil, mocks.UserRepository, nil, mocks.Crypto, nil)
		createdUser, err := usecase.User.SignUp(context.Background(), email, password)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, createdUser)
	})

	t.Run("User already exists", func(t *testing.T) {
		mocks := NewSuite(t)

		var (
			expectedUser                      = models.UserDTO{}
			expectedError                     = customError.ErrUserExists
			getUserByEmailResponse            = models.User{ID: 1, Email: email}
			getUserByEmailResponseError error = nil
		)

		mocks.UserRepository.
			EXPECT().
			GetUserByEmail(gomock.Any(), email).
			Return(getUserByEmailResponse, getUserByEmailResponseError)

		usecase := usecase.New(config.Config{}, nil, mocks.UserRepository, nil, mocks.Crypto, nil)
		createdUser, err := usecase.User.SignUp(context.Background(), email, password)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, expectedUser, createdUser)
	})
}
