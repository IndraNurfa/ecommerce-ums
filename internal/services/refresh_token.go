package services

import (
	"context"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"time"

	"github.com/pkg/errors"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	var (
		now = time.Now()
	)

	resp := models.RefreshTokenResponse{}

	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, tokenClaim.Email, "token", time.Now())
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate new token")
	}

	err = s.UserRepo.UpdateTokenByRefreshToken(ctx, token, refreshToken, now.Add(helpers.MapTypeToken["token"]), now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update new refresh token")
	}
	resp.Token = token
	return resp, nil
}
