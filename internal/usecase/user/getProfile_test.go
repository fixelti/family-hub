package user_test

import (
	"context"
	"testing"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetProfile(t *testing.T) {
	t.Parallel()
	t.Helper()

	const (
		email         = "test@test.com"
		userID   uint = 1
		password      = "test_password"
	)

	t.Run("Success", func(t *testing.T) {
		mocks := NewSuite(t)

		var (
			getUserByIDResponse            = models.User{ID: userID, Password: password, Email: email}
			getUserByIDResponseError error = nil

			getUserServicesResponse = []models.DiskSpaceAllocationService{
				{ID: 1, UserID: userID, Name: "test name", DiskSize: 1000, Status: models.Active},
				{ID: 2, UserID: userID, Name: "test name2", DiskSize: 1250, Status: models.Active},
			}
			getUserServicesResponseError error = nil

			expectedUserProfile = models.UserProfile{Email: email, DiskSpaceAllocationService: getUserServicesResponse}
		)

		mocks.UserRepository.
			EXPECT().
			GetUserByID(gomock.Any(), userID).
			Return(getUserByIDResponse, getUserByIDResponseError)

		mocks.DiskSASRepository.
			EXPECT().
			GetUserServices(gomock.Any(), userID).
			Return(getUserServicesResponse, getUserServicesResponseError)

		usecase := usecase.New(config.Config{}, nil, mocks.UserRepository, mocks.DiskSASRepository, mocks.Crypto, nil)
		userProfile, err := usecase.User.GetProfile(context.Background(), userID)
		assert.NoError(t, err)
		assert.Equal(t, expectedUserProfile, userProfile)
	})
}
