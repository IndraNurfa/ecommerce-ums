package cmd

import (
	"ecommerce-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			log.Println("authorization empty")
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

		id, err := uuid.Parse(claim.ID)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "failed to verify token", nil)
		}

		tokenData, err := d.UserRepository.GetUserSessionById(c.Request().Context(), id)
		if err != nil {
			log.Println("failed to get user session on DB: ", err)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		if helpers.GenerateHash(auth) != tokenData.Token {
			log.Println("token invalid: ", id)
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

		claim, err := helpers.ValidateToken(c.Request().Context(), auth)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		if time.Now().Unix() > claim.ExpiresAt.Unix() {
			log.Println("jwt token is expired: ", claim.ExpiresAt)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		id, err := uuid.Parse(claim.ID)
		if err != nil {
			log.Println(err)
			return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "failed to verify token", nil)
		}

		tokenData, err := d.UserRepository.GetUserSessionById(c.Request().Context(), id)
		if err != nil {
			log.Println("failed to get user session on DB: ", err)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		if helpers.GenerateHash(auth) != tokenData.RefreshToken {
			log.Println("token invalid: ", id)
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		c.Set("token", *claim)

		return next(c)
	}
}
