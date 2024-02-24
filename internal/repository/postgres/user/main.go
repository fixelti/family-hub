package user

import (
	"context"

	_ "github.com/fixelti/family-hub/internal"
	customError "github.com/fixelti/family-hub/internal"
	"github.com/fixelti/family-hub/internal/common/models"
	queryes "github.com/fixelti/family-hub/internal/repository/postgres/user/internal"
	"github.com/fixelti/family-hub/lib/database/postgres"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(ctx context.Context, email, password string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type repository struct {
	db     postgres.Database
	logger *zap.Logger
}

func New(db postgres.Database, logger *zap.Logger) UserRepository {
	return repository{db: db, logger: logger}
}

// Create создание пользователя
func (user repository) Create(ctx context.Context, email, password string) (models.User, error) {
	tx, err := user.db.Begin(ctx)
	if err != nil {
		user.logger.Error("failed to begin transaction: ", zap.Error(err))
		return models.User{}, customError.ErrBeginTransaction
	}

	res, err := tx.Query(
		ctx,
		queryes.Create,
		email,
		password)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			user.logger.Error("failed to rollback transaction: ", zap.Error(err))
		}
		user.logger.Error("failed to make request : ", zap.Error(err))
		return models.User{}, customError.ErrMakeQuery
	}

	createdUser, err := postgres.ScanInStruct[models.User](res)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			user.logger.Error("failed to rollback transaction: ", zap.Error(err))
		}
		user.logger.Error("failed to scan in struct: ", zap.Error(err))
		return models.User{}, customError.ErrScanInSctruct
	}
	if err := tx.Commit(ctx); err != nil {
		user.logger.Error("failed to commit transaction: ", zap.Error(err))
	}

	return *createdUser, nil
}

// GetUserByEmail возвращает пользователя по его email адрессу
func (user repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	tx, err := user.db.Begin(ctx)
	if err != nil {
		user.logger.Error("failed to begin transaction: ", zap.Error(err))
		return models.User{}, customError.ErrBeginTransaction
	}

	res, err := tx.Query(
		ctx,
		queryes.GetByEmail,
		email)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			user.logger.Error("failed to rollback transaction: ", zap.Error(err))
		}
		user.logger.Error("failed to make request : ", zap.Error(err))
		return models.User{}, customError.ErrMakeQuery
	}

	foundUser, err := postgres.ScanInStruct[models.User](res)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			user.logger.Error("failed to rollback transaction: ", zap.Error(err))
		}
		user.logger.Error("failed to scan in struct: ", zap.Error(err))
		return models.User{}, customError.ErrScanInSctruct
	}

	if err := tx.Commit(ctx); err != nil {
		user.logger.Error("failed to commit transaction: ", zap.Error(err))
	}

	return *foundUser, nil
}
