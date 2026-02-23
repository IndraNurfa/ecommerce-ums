package interfaces

import (
	"context"
	"ecommerce-ums/internal/models"

	"github.com/labstack/echo/v5"
)

type IUserAPI interface {
	RegisterUser(e *echo.Context) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, req *models.User) (*models.User, error)
}

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}
