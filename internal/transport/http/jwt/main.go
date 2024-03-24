package jwt

import (
	"github.com/fixelti/family-hub/lib/jwt"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	jwt    jwt.JWT
}

func New(logger *zap.Logger, jwt jwt.JWT) Handler {
	return Handler{logger: logger, jwt: jwt}
}
