package services

import (
	"context"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"time"
)

type HealthService struct {
	HealthRepo interfaces.IHealthRepository
}

func (s *HealthService) CheckHealthConnection(ctx context.Context) (*models.Health, error) {
	err := s.HealthRepo.CheckDatabaseConnection(ctx)
	if err != nil {
		return nil, err
	}
	return &models.Health{
		Message: "Service is healthy",
		Time:    time.Now(),
	}, nil
}
