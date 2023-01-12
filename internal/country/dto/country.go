package dto

import (
	"github.com/gofrs/uuid"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/model"
)

type CountryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (request *CountryRequest) ToModel() *model.Country {
	return &model.Country{
		Name: request.Name,
	}
}

type CountryResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewCountryResponse(model *model.Country) *CountryResponse {
	return &CountryResponse{
		Id:   model.ID,
		Name: model.Name,
	}
}
