package user

import (
	"context"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres/user"
)

type Usecase interface {
	SignUp(ctx context.Context, email, password string) (uint, error)
	SignIn(ctx context.Context, email, password string) (models.Tokens, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (accessToken string, err error)
}

type userUsecase struct {
	db     user.UserRepository
	config config.Config
}

func New(db user.UserRepository, config config.Config) Usecase {
	return userUsecase{
		db:     db,
		config: config,
	}
}
