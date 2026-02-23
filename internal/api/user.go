package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserAPI struct {
	UserService interfaces.IUserService
}

func (api *UserAPI) RegisterUser(e *echo.Context) error {
	var (
		req *models.User
		log = helpers.Logger
	)

	if err := e.Bind(req); err != nil {
		log.Error("failed to parese request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := e.Validate(req); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.UserService.RegisterUser(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to register : ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
