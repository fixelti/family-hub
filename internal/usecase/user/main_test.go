package user_test

import (
	"testing"

	mock_postgres "github.com/fixelti/family-hub/internal/repository/postgres/mocks"
	mock_user "github.com/fixelti/family-hub/internal/repository/postgres/user/mocks"
	mock_crypto "github.com/fixelti/family-hub/lib/crypto/mocks"
	"go.uber.org/mock/gomock"
)

type Mocks struct {
	RepositoryManager *mock_postgres.MockRepository // TODO: возможно надо удалить
	UserRepository    *mock_user.MockUserRepository
	Crypto            *mock_crypto.MockCrypto
}

func NewSuite(t *testing.T) Mocks {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return Mocks{
		UserRepository:    mock_user.NewMockUserRepository(ctrl),
		RepositoryManager: mock_postgres.NewMockRepository(ctrl),
		Crypto:            mock_crypto.NewMockCrypto(ctrl),
	}
}
