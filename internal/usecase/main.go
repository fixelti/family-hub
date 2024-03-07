package usecase

import (
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres"
	"github.com/fixelti/family-hub/internal/usecase/user"
	"go.uber.org/zap"
)

type UsecaseManager struct {
	User                       user.Usecase
}

func New(config config.Config, logger *zap.Logger, repositoryManager postgres.RepositoryManager) UsecaseManager {
	return UsecaseManager{
		User: user.New(repositoryManager.User, repositoryManager.DiskSpaceAllocationService, config, logger),
	}
}
