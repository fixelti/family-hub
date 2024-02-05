package user

import "github.com/fixelti/family-hub/internal/usecase/user"

type Handler struct {
	Usecase user.Usecase
}

func New(userUsecase user.Usecase) Handler {
	return Handler{Usecase: userUsecase}
}