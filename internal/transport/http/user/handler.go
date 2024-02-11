package user

import (
	"log"
	"net/http"
	"strings"

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

func (handler Handler) SingIn(c echo.Context) error {
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

	tokens, err := handler.Usecase.SignIn(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		if err.Error() == ErrUserExists.Error() {
			c.JSON(http.StatusConflict, map[string]string{"error": "user exists"})
			return err
		}
		log.Printf("signup -> failed to signup usecase: %s\n", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(http.StatusOK, tokens)
}

func (handler Handler) RefreshAccessToken(c echo.Context) error {
	refreshToken := c.Request().Header.Get("Authorization")
	bearerToken := strings.Split(refreshToken, "Bearer ")
	refreshToken = bearerToken[1]

	accessToken, err := handler.Usecase.RefreshAccessToken(c.Request().Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return err
	}

	c.JSON(http.StatusOK, accessToken)
	return nil
}