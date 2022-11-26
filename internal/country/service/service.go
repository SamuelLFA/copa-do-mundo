package service

import (
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

func (service *Service) CreateCountry(request *dto.CountryRequest) dto.CountryResponse {
	model := request.ToModel()
	service.repository.Save(model)
	return dto.NewCountryResponse(model)
}
