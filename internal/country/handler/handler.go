package handler

import (
	"encoding/json"
	"net/http"

	"github.com/samuellfa/copa-do-mundo-golang/internal/country/dto"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/service"
	"github.com/samuellfa/copa-do-mundo-golang/internal/shared"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// ShowAccount 	godoc
// @Summary     Process a new country
// @Description Endpoint to create a new country
// @Accept      json
// @Produce     json
// @Param       body body     CountryRequest true "Country request payload"
// @Success     201  {object} string
// @Success     400  {object} string
// @Router      /country [post]
func (h *Handler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	var request dto.CountryRequest
	if err := dto.ValidateInput(r.Body, &request); err != nil {
		customError := err.(*shared.RequestError)
		w.WriteHeader(customError.StatusCode)
		w.Write([]byte(customError.Err.Error()))
		return
	}

	response, err := createCountry(&request, *h.service)
	if err != nil {
		customError := err.(*shared.RequestError)
		w.WriteHeader(customError.StatusCode)
		w.Write([]byte(customError.Err.Error()))
		return
	}
	text, _ := json.Marshal(response)
	w.WriteHeader(http.StatusCreated)
	w.Write(text)
}

func createCountry(request *dto.CountryRequest, service service.Service) (*dto.CountryResponse, error) {
	return service.CreateCountry(request)
}
