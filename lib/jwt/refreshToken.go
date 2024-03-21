package jwt

import (
	"context"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	jwtLib "github.com/golang-jwt/jwt"
)

func (j jwt) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims := jwtLib.MapClaims{}
	tkn, err := jwtLib.ParseWithClaims(refreshToken, &claims, func(token *jwtLib.Token) (any, error) {
		return []byte(j.RefreshKey), nil
	})
	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", customError.ErrTokenIsNotValid
	}

	tokens, err := j.GenerateTokens(uint(claims["id"].(float64)))
	if err != nil {
		return tokens.Access, err
	}

	return tokens.Access, nil
}
