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

	v1 := e.Group("/api/v1")

	// User routes
	v1.POST("/users/register", d.UserAPI.RegisterUser)
	v1.POST("/users/register/admin", d.UserAPI.RegisterAdmin)
	v1.POST("/users/login", d.UserAPI.LoginUser)
	v1.POST("/users/login/admin", d.UserAPI.LoginAdmin)

	if err := e.Start(":" + helpers.GetEnv("PORT", "9000")); err != nil {
		helpers.Logger.Error("failed to start server: ", err)
	}
}

type Dependency struct {
	HealthAPI      *api.HealthAPI
	UserRepository interfaces.IUserRepository
	UserAPI        interfaces.IUserAPI
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

	return Dependency{
		HealthAPI:      healthAPI,
		UserRepository: userRepo,
		UserAPI:        userAPI,
	}
}
