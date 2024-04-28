package repository

import (
	"gorm.io/gorm"
)

const TableUser = "user"
const TableContact = "contact"

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
