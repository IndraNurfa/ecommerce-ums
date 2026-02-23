package interfaces

import (
	"context"
	"ecommerce-ums/internal/models"
)

type IHealthService interface {
	CheckHealthConnection(ctx context.Context) (*models.Health, error)
}

type IHealthRepository interface {
	CheckDatabaseConnection(ctx context.Context) error
}
