package usecase

import (
	"github.com/fixelti/family-hub/internal/usecase/user"
	userRepo "github.com/fixelti/family-hub/internal/repository/postgres/user"
)

type UsecaseManager struct {
	User user.Usecase
}


func New(userRepo userRepo.UserRepository) UsecaseManager {
	return UsecaseManager{
		User: user.New(userRepo),
	}
}
