package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/repository"
	"ecommerce-ums/internal/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func ServeHTTP() {
	d := dependencyInject()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())
	e.GET("/health", d.HealthAPI.Health)

	v1 := e.Group("/api/v1")

	// User routes
	v1.POST("/users/register", d.UserAPI.RegisterUser)

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
