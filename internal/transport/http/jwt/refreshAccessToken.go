package jwt

import (
	internalHttp "net/http"

	customError "github.com/fixelti/family-hub/internal/common/errors"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (handler Handler) RefreshAccessToken(c echo.Context) error {
	refreshToken := c.Param("refresh_token")
	if len(refreshToken) == 0 {
		c.JSON(internalHttp.StatusUnauthorized, nil)
		return customError.ErrInvalidCredentials
	}

	accessToken, err := handler.jwt.RefreshToken(c.Request().Context(), refreshToken)
	if err != nil {
		handler.logger.Error("failed to refresh access token", zap.Error(err))
		c.JSON(internalHttp.StatusInternalServerError, nil)
		return customError.ErrInternal
	}

	c.JSON(internalHttp.StatusOK, echo.Map{"access_token": accessToken})
	return nil
}
