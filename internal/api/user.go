package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserAPI struct {
	UserService interfaces.IUserService
}

func (api *UserAPI) RegisterUser(e *echo.Context) error {
	req := &models.User{}
	log := helpers.Logger

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
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

func (api *UserAPI) RegisterAdmin(e *echo.Context) error {
	req := &models.User{}
	log := helpers.Logger

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.UserService.RegisterAdmin(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to register admin : ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *UserAPI) LoginUser(e *echo.Context) error {
	req := &models.LoginRequest{}
	log := helpers.Logger

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.UserService.Login(e.Request().Context(), req, string(models.RoleUser))
	if err != nil {
		log.Error("failed to login : ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *UserAPI) LoginAdmin(e *echo.Context) error {
	req := &models.LoginRequest{}
	log := helpers.Logger

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.UserService.Login(e.Request().Context(), req, string(models.RoleAdmin))
	if err != nil {
		log.Error("failed to login as admin : ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *UserAPI) GetProfile(e *echo.Context) error {
	log := helpers.Logger

	token := e.Get("token")
	tokenClaim, ok := token.(helpers.ClaimToken)
	fmt.Println("ok", ok)
	if !ok {
		log.Error("failed to fetch token")
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}

	resp, err := api.UserService.GetProfile(e.Request().Context(), tokenClaim.Username)
	if err != nil {
		log.Error("failed to get profile : ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
