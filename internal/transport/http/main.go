package httpTransport

import (
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/transport/http/user"
	userHandler "github.com/fixelti/family-hub/internal/transport/http/user"
	userUsecase "github.com/fixelti/family-hub/internal/usecase/user"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Http struct {
	echo      *echo.Echo
	validator *echo.Validator
	config    config.Config
	logger    *zap.Logger

	userHandler user.Handler
}

type CustomValidator struct {
	validator *validator.Validate
}

func New(userUsecase userUsecase.Usecase, 
	config config.Config, 
	logger *zap.Logger) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	http := Http{
		logger: logger,
		config:      config,
		echo:        e,
		userHandler: userHandler.New(userUsecase),
	}
	http.routing()
	return http.echo

}

func (http Http) routing() {
	user := http.echo.Group("/user")
	user.POST("/signup", http.userHandler.SingUp)
	user.POST("/signin", http.userHandler.SingIn)
	user.POST("/refresh-access-token", http.userHandler.RefreshAccessToken)

}

func (cv *CustomValidator) Validate(data interface{}) error {
	if err := cv.validator.Struct(data); err != nil {
		return err
	}
	return nil
}
