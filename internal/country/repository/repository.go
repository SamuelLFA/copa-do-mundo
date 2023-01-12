package repository

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/model"
	"github.com/samuellfa/copa-do-mundo-golang/internal/shared"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *gorm.DB
}

func (repository *Repository) Save(model *model.Country) error {
	uuid := uuid.Must(uuid.NewV4())
	model.ID = uuid

	if result := repository.db.Create(model); result.Error != nil {
		return &shared.RequestError{
			StatusCode: 409,
			Err:        fmt.Errorf("error when save country"),
		}
	}
	return nil
}

func (repository *Repository) GetByName(countryToSave *model.Country) error {
	where := model.Country{Name: countryToSave.Name}
	var countryWithSameName model.Country
	repository.db.First(&countryWithSameName, where)
	if countryWithSameName.ID != uuid.Nil {
		return &shared.RequestError{
			StatusCode: 409,
			Err:        fmt.Errorf("country name already registered"),
		}
	}
	return nil
}
