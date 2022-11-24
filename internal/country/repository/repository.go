package repository

import (
	"github.com/gofrs/uuid"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/model"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *gorm.DB
}

func (repository *Repository) Save(model *model.Country) {
	uuid := uuid.Must(uuid.NewV4())
	model.ID = uuid

	repository.db.Save(model)
}
