package interfaces

import (
	"context"
	"ecommerce-ums/internal/models"

	"github.com/labstack/echo/v5"
)

type IUserAPI interface {
	RegisterUser(e *echo.Context) error
	RegisterAdmin(e *echo.Context) error
	LoginUser(e *echo.Context) error
	LoginAdmin(e *echo.Context) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, req *models.User) (*models.User, error)
	RegisterAdmin(ctx context.Context, req *models.User) (*models.User, error)
	Login(ctx context.Context, req *models.LoginRequest, role string) (models.LoginResponse, error)
}

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserbyUsername(ctx context.Context, username, role string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
}
