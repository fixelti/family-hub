package diskSpaceAllocationService

import (
	"context"
	"errors"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	queryes "github.com/fixelti/family-hub/internal/repository/postgres/diskSpaceAllocationService/internal"
	"github.com/fixelti/family-hub/lib/database/postgres"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type DiskSpaceAllocationServiceRepository interface {
	GetUserServices(ctx context.Context, userID uint) ([]models.DiskSpaceAllocationService, error)
}

type repository struct {
	db     postgres.Database
	logger *zap.Logger
}

func New(db postgres.Database, logger *zap.Logger) DiskSpaceAllocationServiceRepository {
	return repository{
		db:     db,
		logger: logger,
	}
}

// GetUserService возвращает все сервисы конкретного пользователя
func (service repository) GetUserServices(ctx context.Context, userID uint) ([]models.DiskSpaceAllocationService, error) {
	tx, err := service.db.Begin(ctx)
	if err != nil {
		service.logger.Error("failed to begin transaction: ", zap.Error(err))
		return nil, customError.ErrBeginTransaction
	}

	res, err := tx.Query(
		ctx,
		queryes.GetUserServices,
		userID,
	)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			service.logger.Error("failed to rollback transaction: ", zap.Error(err))
		}
		service.logger.Error("failed to make request: ", zap.Error(err))
		return nil, customError.ErrMakeQuery
	}

	diskSpaceAllocationServices, err := postgres.ScanInStruct[[]models.DiskSpaceAllocationService](res)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			service.logger.Error("failed to rollback transaction: ", zap.Error(err))
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		service.logger.Error("failed to scan in struct: ", zap.Error(err))
		return nil, customError.ErrScanInSctruct
	}
	if err := tx.Commit(ctx); err != nil {
		service.logger.Error("failed to commit transaction: ", zap.Error(err))
	}

	return *diskSpaceAllocationServices, nil
}
