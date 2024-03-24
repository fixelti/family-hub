package user_test

import (
	"testing"

	mock_diskSpaceAllocationService "github.com/fixelti/family-hub/internal/repository/postgres/diskSpaceAllocationService/mocks"
	mock_user "github.com/fixelti/family-hub/internal/repository/postgres/user/mocks"
	mock_crypto "github.com/fixelti/family-hub/lib/crypto/mocks"
	"go.uber.org/mock/gomock"
)

type Mocks struct {
	UserRepository    *mock_user.MockUserRepository
	DiskSASRepository *mock_diskSpaceAllocationService.MockDiskSpaceAllocationServiceRepository
	Crypto            *mock_crypto.MockCrypto
}

func NewSuite(t *testing.T) Mocks {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return Mocks{
		UserRepository:    mock_user.NewMockUserRepository(ctrl),
		DiskSASRepository: mock_diskSpaceAllocationService.NewMockDiskSpaceAllocationServiceRepository(ctrl),
		Crypto:            mock_crypto.NewMockCrypto(ctrl),
	}
}
