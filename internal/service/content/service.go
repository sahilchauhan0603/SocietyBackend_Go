package content

import "github.com/sahilchauhan0603/society/internal/repository"

type Service struct {
	repo repository.ContentRepository
}

func NewService(repo repository.ContentRepository) *Service {
	return &Service{repo: repo}
}

func NewDefaultService() *Service {
	return NewService(repository.NewContentRepository())
}
