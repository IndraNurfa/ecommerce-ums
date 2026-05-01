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

func (s *RefreshTokenService) RefreshToken(ctx context.Context, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	var (
		now = time.Now()
	)

	resp := models.RefreshTokenResponse{}

	token, err := helpers.GenerateToken(ctx, tokenClaim.ID, tokenClaim.Subject, "token", time.Now())
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate new token")
	}

	hashNewToken := helpers.GenerateHash(token)

	err = s.UserRepo.UpdateTokenById(ctx, hashNewToken, tokenClaim.ID, now.Add(helpers.MapTypeToken["token"]), now)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update new refresh token")
	}
	resp.Token = token
	return resp, nil
}
