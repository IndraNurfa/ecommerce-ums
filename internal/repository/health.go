package repository

import (
	"context"

	"gorm.io/gorm"
)

type HealthRepository struct {
	DB *gorm.DB
}

func (r *HealthRepository) CheckDatabaseConnection(ctx context.Context) error {
	return r.DB.Exec("SELECT 1").Error
}
