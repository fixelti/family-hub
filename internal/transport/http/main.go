package httpTransport

import (
	"github.com/fixelti/family-hub/internal/transport/http/user"
	userHandler "github.com/fixelti/family-hub/internal/transport/http/user"
	userUsecase "github.com/fixelti/family-hub/internal/usecase/user"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Http struct {
	echo      *echo.Echo
	validator *echo.Validator

	userHandler user.Handler
}

type CustomValidator struct {
	validator *validator.Validate
}

func New(userUsecase userUsecase.Usecase) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	http := Http{
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
