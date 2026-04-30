package interfaces

import (
	"context"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/models"

	"github.com/labstack/echo/v5"
)

type IRefreshTokenHandler interface {
	RefreshToken(e *echo.Context) error
}

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}
