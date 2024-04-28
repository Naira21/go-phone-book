package createService

import (
	"go-phone-book/src/domain/entity/user"
)

type (
	UserCreateRepository interface {
		CreateUser(user *userEntity.User) error
	}

	Service struct {
		repository UserCreateRepository
	}
)

func NewService(repository UserCreateRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(name, surname string) error {
	newUser := userEntity.Create(name, surname)

	return s.repository.CreateUser(newUser)
}
