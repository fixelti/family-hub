package user

import (
	"errors"
	"net/http"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/labstack/echo/v4"
)

func (handler Handler) SingUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return customError.ErrBind
	}

	if err := c.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return customError.ErrBind
	}

	createdUser, err := handler.Usecase.SignUp(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		if errors.Is(err, customError.ErrUserExists) {
			c.JSON(http.StatusConflict, echo.Map{"error": err.Error()})
			return err
		}
		c.JSON(http.StatusInternalServerError, nil)
		return err
	}

	return c.JSON(http.StatusOK, createdUser)
}
