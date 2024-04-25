package createService

import "go-phone-book/src/repository"

type (
	Service struct {
		repository repository.Repository
	}
)

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
