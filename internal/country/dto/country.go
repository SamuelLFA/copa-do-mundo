package dto

import (
	"github.com/gofrs/uuid"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/model"
)

// CountryRequest represents the request payload for creating a new country
// swagger:parameters createCountry
type CountryRequest struct {
	// Name of the country
	//
	// required: true
	Name string `json:"name" validate:"required"`
}

func (request *CountryRequest) ToModel() *model.Country {
	return &model.Country{
		Name: request.Name,
	}
}

// CountryResponse represents the response payload for a country
// swagger:response countryResponse
type CountryResponse struct {
	// The ID of the country
	//
	// example: "123e4567-e89b-12d3-a456-426614174000"
	Id uuid.UUID `json:"id"`
	// The name of the country
	//
	// example: "Brazil"
	Name string `json:"name"`
}

func NewCountryResponse(model *model.Country) *CountryResponse {
	return &CountryResponse{
		Id:   model.ID,
		Name: model.Name,
	}
}
