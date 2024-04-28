package repository

import (
	"gorm.io/gorm"
)

const tableUser = "user"
const tableContact = "contact"

type (
	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
