package repository

import (
	"go-phone-book/src/domain/entity/user"
)

type (
	User struct {
		Name    string `gorm:"name"`
		Surname string `gorm:"surname"`
	}
)

func (r *Repository) CreateUser(user *userModel.User) error {
	newUser := &User{
		Name:    user.Name(),
		Surname: user.Surname(),
	}

	err := r.db.Table(TableUser).Debug().Create(newUser).Error
	if err != nil {
		return err
	}

	return nil
}
