package httpTransport

import (
	"strings"

	internalHttp "net/http"

	customError "github.com/fixelti/family-hub/internal"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (http Http) UserAuthorizationCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		tokenKey := http.config.JWT.TokenKey
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(reqToken, &claims, func(token *jwt.Token) (any, error) {
			return []byte(tokenKey), nil
		})
		if err != nil {
			http.logger.Error("failed to parse jwt claims", zap.Error(err))
			c.JSON(internalHttp.StatusInternalServerError, nil)
			return err
		}
		if !tkn.Valid {
			c.JSON(internalHttp.StatusUnauthorized, nil)
			return customError.ErrInvalidCredentials
		}

		return next(c)
	}
}
