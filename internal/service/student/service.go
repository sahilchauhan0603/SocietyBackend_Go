package student

import "github.com/sahilchauhan0603/society/internal/repository"

type Service struct {
	repo repository.StudentRepository
}

func NewService(repo repository.StudentRepository) *Service {
	return &Service{repo: repo}
}

func NewDefaultService() *Service {
	return NewService(repository.NewStudentRepository())
}
