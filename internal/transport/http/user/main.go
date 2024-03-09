package user

import (
	"errors"
	"net/http"
	"strings"

	internalHttp "net/http"

	"github.com/fixelti/family-hub/internal/common/constants"
	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/usecase/user"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	Usecase user.Usecase
}

func New(userUsecase user.Usecase) Handler {
	return Handler{Usecase: userUsecase}
}

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

func (handler Handler) RefreshAccessToken(c echo.Context) error {
	refreshToken := c.Request().Header.Get("Authorization")
	bearerToken := strings.Split(refreshToken, "Bearer ")
	if len(bearerToken) != 2 {
		c.JSON(internalHttp.StatusForbidden, nil)
		return customError.ErrInvalidCredentials
	}
	refreshToken = bearerToken[1]

	accessToken, err := handler.Usecase.RefreshAccessToken(c.Request().Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusForbidden, nil)
		return err
	}

	c.JSON(http.StatusOK, echo.Map{"access_token": accessToken})
	return nil
}

func (handler Handler) GetUserProfile(c echo.Context) error {
	userID := c.Get(constants.UserIDContextKey)

	userProfile, err := handler.Usecase.GetProfile(c.Request().Context(), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return customError.ErrInternal
	}

	c.JSON(http.StatusOK, userProfile)
	return nil
}
