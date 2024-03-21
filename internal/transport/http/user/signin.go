package user

import (
	"net/http"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/labstack/echo/v4"
)

func (handler Handler) SingIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return err
	}

	if err := c.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return err
	}

	tokens, err := handler.Usecase.SignIn(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return err
	}

	return c.JSON(http.StatusOK, tokens)
}
