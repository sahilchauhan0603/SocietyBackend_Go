package admin

import "github.com/sahilchauhan0603/society/internal/repository"

type Service struct {
	repo repository.AdminRepository
}

func NewService(repo repository.AdminRepository) *Service {
	return &Service{repo: repo}
}

func NewDefaultService() *Service {
	return NewService(repository.NewAdminRepository())
}
