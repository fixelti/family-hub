package user

import (
	"github.com/fixelti/family-hub/internal/usecase/user"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	Usecase user.Usecase
}

func New(userUsecase user.Usecase, logger *zap.Logger) Handler {
	return Handler{Usecase: userUsecase, logger: logger}
}
