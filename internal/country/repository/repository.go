package repository

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/model"
	"github.com/samuellfa/copa-do-mundo-golang/internal/shared"
	"gorm.io/gorm"
)

// New creates a new instance of the repository with a given database connection.
func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Repository holds the database connection.
type Repository struct {
	db *gorm.DB
}

// Save saves a country in the database.
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

// GetByName gets a country by name from the database.
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

// GetById gets a country by ID from the database.
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

// GetAllCountries returns all countries, paginated according to the provided page and limit.
func (repository *Repository) GetAllCountries(page, limit int) ([]model.Country, error) {
	var countries []model.Country
	offset := (page - 1) * limit

	result := repository.db.Offset(offset).Limit(limit).Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}

	// get total count of countries
	var count int64
	if err := repository.db.Model(&model.Country{}).Count(&count).Error; err != nil {
		return nil, err
	}

	return countries, nil
}
