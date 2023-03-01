package service

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/dto"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/repository"
)

type Service struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) CreateCountry(request *dto.CountryRequest) (*dto.CountryResponse, error) {
	model := request.ToModel()
	if countryInDB, err := service.repository.GetByName(model.Name); countryInDB != nil || err != nil {
		if countryInDB != nil {
			return nil, errors.New("country already registered")
		}
		return nil, err
	}
	err := service.repository.Save(model)
	return dto.NewCountryResponse(model), err
}

func (service *Service) GetCountryById(id uuid.UUID) (*dto.CountryResponse, error) {
	model, err := service.repository.GetById(id)
	if err != nil || model == nil {
		return nil, err
	}
	return dto.NewCountryResponse(model), nil
}

func (service *Service) GetAllCountries(page int, limit int) ([]dto.CountryResponse, error) {
	models, err := service.repository.GetAllCountries(page, limit)
	if err != nil {
		return nil, err
	}
	countries := make([]dto.CountryResponse, len(models))
	for i, model := range models {
		countries[i] = *dto.NewCountryResponse(&model)
	}
	return countries, nil
}
