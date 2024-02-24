package postgres

import (
	"github.com/fixelti/family-hub/internal/repository/postgres/user"
	"github.com/fixelti/family-hub/lib/database/postgres"
	"go.uber.org/zap"
)

type RepositoryManager struct {
	User user.UserRepository
}

func New(db postgres.Database, logger *zap.Logger) RepositoryManager {
	return RepositoryManager{
		User: user.New(db, logger),
	}
}
