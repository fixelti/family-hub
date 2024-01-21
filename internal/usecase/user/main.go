package user

import (
	"context"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/repository/postgres/user"
)

type UserUsecase interface {
	SignUp(ctx context.Context, email, password string) (models.UserDTO, error)
}

type userUsecase struct {
	db user.UserRepository
}

func New(db user.UserRepository) UserUsecase {
	return userUsecase{
		db: db,
	}
}
