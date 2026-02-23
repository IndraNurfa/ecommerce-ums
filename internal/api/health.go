package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthAPI struct {
}

func (api *HealthAPI) Health(e *echo.Context) error {
	return e.String(http.StatusOK, "Hello, World!")
}
