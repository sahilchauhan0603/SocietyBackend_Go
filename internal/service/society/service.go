package society

import "github.com/sahilchauhan0603/society/internal/repository"

type Service struct {
	repo repository.SocietyRepository
}

func NewService(repo repository.SocietyRepository) *Service {
	return &Service{repo: repo}
}

func NewDefaultService() *Service {
	return NewService(repository.NewSocietyRepository())
}
