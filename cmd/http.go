package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func ServeHTTP() {
	healthAPI := &api.HealthAPI{}
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.GET("/health", healthAPI.Health)

	if err := e.Start(":" + helpers.GetEnv("PORT", "9000")); err != nil {
		helpers.Logger.Error("failed to start server: ", err)
	}
}
