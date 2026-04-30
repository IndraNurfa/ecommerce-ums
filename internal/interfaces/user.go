package interfaces

import (
	"context"
	"ecommerce-ums/internal/models"
	"time"

	"github.com/labstack/echo/v5"
)

type IUserAPI interface {
	RegisterUser(e *echo.Context) error
	RegisterAdmin(e *echo.Context) error
	LoginUser(e *echo.Context) error
	LoginAdmin(e *echo.Context) error
	GetProfile(e *echo.Context) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, req *models.User) (*models.User, error)
	RegisterAdmin(ctx context.Context, req *models.User) (*models.User, error)
	Login(ctx context.Context, req *models.LoginRequest, role string) (models.LoginResponse, error)
	GetProfile(ctx context.Context, username string) (models.User, error)
}

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserbyUsername(ctx context.Context, username, role string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
	UpdateTokenByRefreshToken(ctx context.Context, token, refresh_token string, tokenExpired, updatedAt time.Time) error
}
