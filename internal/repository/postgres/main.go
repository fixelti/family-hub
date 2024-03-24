package postgres

import (
	"github.com/fixelti/family-hub/internal/repository/postgres/diskSpaceAllocationService"
	"github.com/fixelti/family-hub/internal/repository/postgres/user"
	"github.com/fixelti/family-hub/lib/database/postgres"
	"go.uber.org/zap"
)


type Repository interface {
	user.UserRepository
	diskSpaceAllocationService.DiskSpaceAllocationServiceRepository
}

type RepositoryManager struct {
	User                       user.UserRepository
	DiskSpaceAllocationService diskSpaceAllocationService.DiskSpaceAllocationServiceRepository
}

func New(db postgres.Database, logger *zap.Logger) RepositoryManager {
	return RepositoryManager{
		User:                       user.New(db, logger),
		DiskSpaceAllocationService: diskSpaceAllocationService.New(db, logger),
	}
}
