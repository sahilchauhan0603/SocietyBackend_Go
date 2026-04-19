package repository

import (
	"context"

	database "github.com/sahilchauhan0603/society/internal/database"
)

type HealthRepository interface {
	Ping(ctx context.Context) error
}

type healthRepository struct{}

func NewHealthRepository() HealthRepository {
	return &healthRepository{}
}

func (r *healthRepository) Ping(ctx context.Context) error {
	return database.PingContext(ctx)
}
