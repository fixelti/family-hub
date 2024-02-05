package usecase

import (
	"github.com/fixelti/family-hub/internal/config"
	userRepo "github.com/fixelti/family-hub/internal/repository/postgres/user"
	"github.com/fixelti/family-hub/internal/usecase/user"
)

type UsecaseManager struct {
	User user.Usecase
}

func New(config config.Config, userRepo userRepo.UserRepository) UsecaseManager {
	return UsecaseManager{
		User: user.New(userRepo, config),
	}
}
