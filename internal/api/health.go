package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthAPI struct {
	HealthService interfaces.IHealthService
}

func (api *HealthAPI) Health(e *echo.Context) error {
	var log = helpers.Logger
	resp, err := api.HealthService.CheckHealthConnection(e.Request().Context())
	if err != nil {
		log.Error("Error checking health connection: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}
	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
