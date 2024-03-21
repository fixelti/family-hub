package jwt

import (
	"context"
	"time"

	"github.com/fixelti/family-hub/internal/common/models"
)

//go:generate mockgen -source ./main.go -destination ./mocks/main.go
type JWT interface {
	GenerateTokens(userID uint) (tokens models.Tokens, err error)
	RefreshToken(ctx context.Context, refreshToken string) (accessToken string, err error)
}
type jwt struct {
	AccessKey          string
	RefreshKey         string
	ExpiraAccessToken  time.Duration
	ExpiraRefreshToken time.Duration
}

func New(
	accessKey, refreshKey string,
	expiraAccessToken, expiraRefreshToken time.Duration,
) JWT {
	return jwt{
		AccessKey:          accessKey,
		RefreshKey:         refreshKey,
		ExpiraAccessToken:  expiraAccessToken,
		ExpiraRefreshToken: expiraRefreshToken,
	}
}
