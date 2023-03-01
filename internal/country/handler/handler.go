package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
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

// CreateCountry 	godoc
// @Summary     Process a new country
// @Description Endpoint to create a new country
// @Accept      json
// @Produce     json
// @Param       body body     dto.CountryRequest true "Country request payload"
// @Success     201  {object} dto.CountryResponse
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
		if err.Error() == "country already registered" {
			w.WriteHeader(409)
			w.Write([]byte("country already registered"))
			return
		}

		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}
	text, _ := json.Marshal(response)
	w.WriteHeader(http.StatusCreated)
	w.Write(text)
}

func createCountry(request *dto.CountryRequest, service service.Service) (*dto.CountryResponse, error) {
	return service.CreateCountry(request)
}

// GetCountryByName 	godoc
// @Summary     Get country by name
// @Description Endpoint to retrieve a country by name
// @Accept      json
// @Produce     json
// @Param       country query     string  true    "Country name"
// @Success     200  {object} dto.CountryResponse
// @Success     404  {object} string
// @Router      /country [get]
func (h *Handler) GetCountryByName(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Query().Get("country")

	response, err := getCountryByName(countryName, *h.service)
	if response == nil {
		w.WriteHeader(404)
		w.Write([]byte("not found"))
		return
	}

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	text, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(text)
}

func getCountryByName(name string, service service.Service) (*dto.CountryResponse, error) {
	return service.GetCountryByName(name)
}

// GetCountryById   godoc
// @Summary     Get country by ID
// @Description Endpoint to retrieve a country by ID
// @Accept      json
// @Produce     json
// @Param       id path     string  true    "Country ID"
// @Success     200  {object} dto.CountryResponse
// @Success     404  {object} string
// @Router      /country/{id} [get]
func (h *Handler) GetCountryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Verify if the id is a valid UUID
	countryID, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	response, err := getCountryById(countryID, *h.service)
	if response == nil {
		w.WriteHeader(404)
		w.Write([]byte("not found"))
		return
	}

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	text, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(text)
}

func getCountryById(id uuid.UUID, service service.Service) (*dto.CountryResponse, error) {
	return service.GetCountryById(id)
}
