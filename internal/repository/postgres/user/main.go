package user

import (
	"context"
	"log"

	"github.com/fixelti/family-hub/internal/common/models"
	queryes "github.com/fixelti/family-hub/internal/repository/postgres/user/internal"
	"github.com/fixelti/family-hub/lib/database/postgres"
)

type UserRepository interface {
	Create(ctx context.Context, email, password string) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
}

type repository struct {
	db postgres.Database
}

func New(db postgres.Database) UserRepository {
	return repository{db: db}
}

func (user repository) Create(ctx context.Context, email, password string) (models.User, error) {
	tx, err := user.db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to begin transaction: %s", err)
		return models.User{}, err
	}

	rows, err := tx.Query(
		ctx,
		queryes.Create,
		email,
		password)

	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to query: %s", err)
		return models.User{}, err
	}

	result, err := postgres.ScanInStruct[models.User](rows)
	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to collect extract: %s", err)
		return models.User{}, err
	}

	return *result, nil
}

func (user repository) GetByEmail(ctx context.Context, email string) (models.User, error) {
	tx, err := user.db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to begin transaction: %s", err)
		return models.User{}, err
	}

	rows, err := tx.Query(
		ctx,
		queryes.GetByEmail,
		email)

	if err != nil {
		if err.Error() != postgres.ErrNotFound.Error() {
			_ = tx.Rollback(ctx)
			log.Printf("failed to query: %s", err)
			return models.User{}, err
		}
	}

	result, err := postgres.ScanInStruct[models.User](rows)
	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to collect extract: %s", err)
		return models.User{}, err
	}

	return *result, nil
}
