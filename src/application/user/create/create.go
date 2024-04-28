package create

import (
	"go-phone-book/src/domain/entity/user"
)

type (
	userCreateRepository interface {
		Create(user *userModel.User) error
	}

	UserCreateService struct {
		repository userCreateRepository
	}
)

func (s *UserCreateService) Create(name, surname string) error {
	newUser := userModel.Create(name, surname)

	return s.repository.Create(newUser)
}
