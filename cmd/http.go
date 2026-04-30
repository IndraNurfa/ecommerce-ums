package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/repository"
	"ecommerce-ums/internal/services"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func ServeHTTP() {
	d := dependencyInject()
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.GET("/health", d.HealthAPI.Health)

	v1 := e.Group("/api/v1/users")

	// User routes
	v1.POST("/register", d.UserAPI.RegisterUser)
	v1.POST("/register/admin", d.UserAPI.RegisterAdmin)
	v1.POST("/login", d.UserAPI.LoginUser)
	v1.POST("/login/admin", d.UserAPI.LoginAdmin)

	v1.GET("/profile", d.UserAPI.GetProfile, d.MiddlewareValidateAuth)

	v1.DELETE("/logout", d.UserAPI.Logout, d.MiddlewareValidateAuth)

	v1.PUT("/refresh-token", d.RefreshTokenAPI.RefreshToken, d.MiddlewareRefreshToken)

	if err := e.Start(":" + helpers.GetEnv("PORT", "9000")); err != nil {
		helpers.Logger.Error("failed to start server: ", err)
	}
}

type Dependency struct {
	HealthAPI *api.HealthAPI

	UserRepository interfaces.IUserRepository
	UserAPI        interfaces.IUserAPI

	RefreshTokenAPI interfaces.IRefreshTokenHandler
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	userSvc := &services.UserService{
		UserRepo: userRepo,
	}
	userAPI := &api.UserAPI{
		UserService: userSvc,
	}

	healthRepo := &repository.HealthRepository{
		DB: helpers.DB,
	}
	healthSvc := &services.HealthService{
		HealthRepo: healthRepo,
	}
	healthAPI := &api.HealthAPI{
		HealthService: healthSvc,
	}

	refreshTokensvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}

	refrehTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokensvc,
	}

	return Dependency{
		HealthAPI:       healthAPI,
		UserRepository:  userRepo,
		UserAPI:         userAPI,
		RefreshTokenAPI: refrehTokenAPI,
	}
}
