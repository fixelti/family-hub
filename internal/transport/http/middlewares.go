package httpTransport

import (
	"strings"

	internalHttp "net/http"

	"github.com/fixelti/family-hub/internal/common/constants"
	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (http Http) UserAuthorizationCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(internalHttp.StatusUnauthorized, nil)
			return customError.ErrInvalidCredentials
		}
		reqToken = splitToken[1]

		tokenKey := http.config.JWT.TokenKey
		claims := models.Payload{}
		tkn, err := jwt.ParseWithClaims(reqToken, &claims, func(token *jwt.Token) (any, error) {
			return []byte(tokenKey), nil
		})
		if err != nil {
			c.JSON(internalHttp.StatusUnauthorized, nil)
			return err
		}

		if !tkn.Valid {
			c.JSON(internalHttp.StatusUnauthorized, nil)
			return customError.ErrInvalidCredentials
		}

		c.Set(constants.UserIDContextKey, claims.UserID)
		return next(c)
	}
}
