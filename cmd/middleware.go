package cmd

import (
	"ecommerce-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			log.Println("authorization empty")
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		_, err := d.UserRepository.GetUserSessionByToken(c.Request().Context(), auth)
		if err != nil {
			log.Println("failed to get user session on DB: ", err)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		claim, err := helpers.ValidateToken(c.Request().Context(), auth)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("jwt token is expired: ", claim.ExpiresAt)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		c.Set("token", *claim)

		return next(c)
	}

}

func (d *Dependency) MiddlewareRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized empty", nil)
		}

		_, err := d.UserRepository.GetUserSessionByRefreshToken(c.Request().Context(), auth)
		if err != nil {
			log.Println("failed to get user session on DB: ", err)
			helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		claim, err := helpers.ValidateToken(c.Request().Context(), auth)
		if err != nil {
			log.Println(err)
			helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println(err)
			helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		c.Set("token", *claim)

		return next(c)
	}
}
