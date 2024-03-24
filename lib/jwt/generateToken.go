package jwt

import (
	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	jwtLib "github.com/golang-jwt/jwt"
)

func (j jwt) GenerateTokens(userID uint) (tokens models.Tokens, err error) {
	claims := jwtLib.MapClaims{
		"id":       userID,
		"expirate": j.ExpiraAccessToken,
	}

	tokens.Access, err = generateToken(j.AccessKey, claims)
	if err != nil {
		return tokens, err
	}

	claims["expirate"] = j.ExpiraRefreshToken
	tokens.Refresh, err = generateToken(j.RefreshKey, claims)
	return tokens, err
}

func generateToken(tokenKey string, claims jwtLib.MapClaims) (string, error) {
	tokenWithClaims := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	if tokenWithClaims == nil {
		return "", customError.ErrTokenWithClaimsIsNil
	}

	return tokenWithClaims.SignedString([]byte(tokenKey))
}
