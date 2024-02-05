package user

import (
	"log"
	"net/http"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/labstack/echo/v4"
)

func (handler Handler) SingUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Printf("signup -> failed to bind: %s\n", err)
		c.JSON(http.StatusBadRequest, map[string]string{"error": ErrBind.Error()})
		return err
	}

	if err := c.Validate(user); err != nil {
		log.Printf("signup -> failed to validate: %s\n", err)
		c.JSON(http.StatusBadRequest, map[string]string{"error": ErrValidation.Error()})
		return err
	}

	userID, err := handler.Usecase.SignUp(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		if err.Error() == ErrUserExists.Error() {
			c.JSON(http.StatusConflict, map[string]string{"error": "user exists"})
			return err
		}
		log.Printf("signup -> failed to signup usecase: %s\n", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(http.StatusOK, map[string]uint{"user_id": userID})
}
