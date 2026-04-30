package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"net/http"

	"github.com/labstack/echo/v5"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(e *echo.Context) error {
	var (
		log = helpers.Logger
	)

	refreshToken := e.Request().Header.Get("Authorization")
	token := e.Get("token")
	tokenClaim, ok := token.(helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse claim token")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}

	resp, err := api.RefreshTokenService.RefreshToken(e.Request().Context(), refreshToken, tokenClaim)
	if err != nil {
		log.Error("failed on logout service: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
