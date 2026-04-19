package service

import (
	"context"
	"time"

	"github.com/sahilchauhan0603/society/internal/repository"
)

type HealthService interface {
	Check(ctx context.Context) HealthStatus
}

type HealthStatus struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Checks    map[string]string `json:"checks"`
}

type healthService struct {
	repo repository.HealthRepository
}

func NewHealthService(repo repository.HealthRepository) HealthService {
	return &healthService{repo: repo}
}

func (s *healthService) Check(ctx context.Context) HealthStatus {
	status := HealthStatus{
		Status:    "ok",
		Timestamp: time.Now().UTC(),
		Checks: map[string]string{
			"database": "ok",
		},
	}

	if err := s.repo.Ping(ctx); err != nil {
		status.Status = "degraded"
		status.Checks["database"] = "down"
	}

	return status
}
