package dto

import "github.com/samuellfa/copa-do-mundo-golang/internal/country/model"

type CountryRequest struct {
	Name string `json:"name"`
}

func (request *CountryRequest) ToModel() *model.Country {
	return &model.Country{
		Name: request.Name,
	}
}
