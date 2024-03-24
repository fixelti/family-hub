package user

import (
	"net/http"

	"github.com/fixelti/family-hub/internal/common/constants"
	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/labstack/echo/v4"
)

func (handler Handler) GetProfile(c echo.Context) error {
	userID := c.Get(constants.UserIDContextKey)

	userProfile, err := handler.Usecase.GetProfile(c.Request().Context(), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return customError.ErrInternal
	}

	c.JSON(http.StatusOK, userProfile)
	return nil
}
