package httpTransport

import (
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/transport/http/user"
	userHandler "github.com/fixelti/family-hub/internal/transport/http/user"
	userUsecase "github.com/fixelti/family-hub/internal/usecase/user"
	"github.com/fixelti/family-hub/lib/jwt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	jwtHandler "github.com/fixelti/family-hub/internal/transport/http/jwt"
)

type Http struct {
	echo      *echo.Echo
	validator *echo.Validator
	config    config.Config
	logger    *zap.Logger

	userHandler user.Handler
	jwtHandler jwtHandler.Handler
}

type CustomValidator struct {
	validator *validator.Validate
}

func New(userUsecase userUsecase.Usecase,
	config config.Config,
	logger *zap.Logger,
	jwt jwt.JWT) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	http := Http{
		logger:      logger,
		config:      config,
		echo:        e,
		userHandler: userHandler.New(userUsecase, logger),
		jwtHandler: jwtHandler.New(logger, jwt),
	}
	http.routing()
	return http.echo

}

func (cv *CustomValidator) Validate(data interface{}) error {
	if err := cv.validator.Struct(data); err != nil {
		return err
	}
	return nil
}
