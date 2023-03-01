package repository

import (
	"errors"

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
			Err:        errors.New("error when save country"),
		}
	}
	return nil
}

func (repository *Repository) GetByName(name string) (*model.Country, error) {
	var country model.Country
	if result := repository.db.Where("name = ?", name).First(&country); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.New("error when getting country")
	}
	return &country, nil
}

func (repository *Repository) GetById(id uuid.UUID) (*model.Country, error) {
	var country model.Country
	if result := repository.db.First(&country, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.New("error when getting country")
	}
	return &country, nil
}
