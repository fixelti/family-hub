package usecase

import (
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres/diskSpaceAllocationService"
	storageUser "github.com/fixelti/family-hub/internal/repository/postgres/user"
	"github.com/fixelti/family-hub/internal/usecase/user"
	"github.com/fixelti/family-hub/lib/crypto"
	"github.com/fixelti/family-hub/lib/jwt"
	"go.uber.org/zap"
)

type UsecaseManager struct {
	User user.Usecase
}

func New(
	config config.Config,
	logger *zap.Logger,
	repositoryUser storageUser.UserRepository,
	repositoryDisk diskSpaceAllocationService.DiskSpaceAllocationServiceRepository,
	crypto crypto.Crypto,
	jwtLib jwt.JWT) UsecaseManager {
	return UsecaseManager{
		User: user.New(repositoryUser, repositoryDisk, config, logger, crypto, jwtLib),
	}
}
