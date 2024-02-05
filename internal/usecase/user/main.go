package user

import (
	"context"

	"github.com/fixelti/family-hub/internal/repository/postgres/user"
)

type Usecase interface {
	SignUp(ctx context.Context, email, password string) (uint, error)
}

type userUsecase struct {
	db user.UserRepository
}

func New(db user.UserRepository) Usecase {
	return userUsecase{
		db: db,
	}
}
