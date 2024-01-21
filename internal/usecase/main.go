package usecase

import (
	"github.com/fixelti/family-hub/internal/usecase/user"
	userRepo "github.com/fixelti/family-hub/internal/repository/postgres/user"
)

type UsecaseManager struct {
	userUsecase user.UserUsecase
}


func New(userRepo userRepo.UserRepository) UsecaseManager {
	return UsecaseManager{
		userUsecase: user.New(userRepo),
	}
}
