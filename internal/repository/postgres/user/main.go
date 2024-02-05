package user

import (
	"context"
	"log"

	"github.com/fixelti/family-hub/internal/common/models"
	queryes "github.com/fixelti/family-hub/internal/repository/postgres/user/internal"
	"github.com/fixelti/family-hub/lib/database/postgres"
)

type UserRepository interface {
	Create(ctx context.Context, email, password string) (uint, error)
	GetIDByEmail(ctx context.Context, email string) (uint, error)
	GetIDAndPasswordrByEmail(ctx context.Context, email string) (models.User, error)
}

type repository struct {
	db postgres.Database
}

func New(db postgres.Database) UserRepository {
	return repository{db: db}
}

// Create создание пользователя
func (user repository) Create(ctx context.Context, email, password string) (uint, error) {
	var u models.User
	tx, err := user.db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to begin transaction: %s", err)
		return u.ID, err
	}

	err = tx.QueryRow(
		ctx,
		queryes.Create,
		email,
		password).Scan(&u.ID)

	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to query: %s", err)
		return u.ID, err
	}
	tx.Commit(ctx)

	return u.ID, nil
}

// GetByEmail получение пользовательского id по его email
func (user repository) GetIDByEmail(ctx context.Context, email string) (uint, error) {
	var userID uint
	tx, err := user.db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to begin transaction: %s", err)
		return userID, err
	}

	rows, err := tx.Query(
		ctx,
		queryes.GetByEmail,
		email)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&userID); err != nil {
			_ = tx.Rollback(ctx)
			log.Printf("failed to scan: %s", err)
			return userID, err
		}
	}

	if err != nil {
		_ = tx.Rollback(ctx)
		log.Printf("failed to query: %s", err)
		return userID, err
	}

	tx.Commit(ctx)

	return userID, nil
}

func (user repository) GetIDAndPasswordrByEmail(ctx context.Context, email string) (models.User, error) {
	var createdUser models.User
	tx, err := user.db.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		log.Printf("failed to start transaction in GetUserByEmail: %s", err)
		return models.User{}, err
	}

	tx.QueryRow(
		ctx,
		queryes.GetUserIDAndPasswordByEmail,
		email,
	).Scan(&createdUser.ID, &createdUser.Password)

	if createdUser.ID == 0 {
		tx.Rollback(ctx)
		return models.User{}, ErrUserNotFound
	}

	tx.Commit(ctx)

	return createdUser, nil
}
