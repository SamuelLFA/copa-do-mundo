package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
// @Router      /countries [post]
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func createCountry(request *dto.CountryRequest, service service.Service) (*dto.CountryResponse, error) {
	return service.CreateCountry(request)
}

// GetCountryById   godoc
// @Summary     Get country by ID
// @Description Endpoint to retrieve a country by ID
// @Accept      json
// @Produce     json
// @Param       id path     string  true    "Country ID"
// @Success     200  {object} dto.CountryResponse
// @Success     404  {object} string
// @Router      /countries/{id} [get]
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
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getCountryById(id uuid.UUID, service service.Service) (*dto.CountryResponse, error) {
	return service.GetCountryById(id)
}

// GetAllCountries godoc
// @Summary     Get all countries
// @Description Endpoint to retrieve all countries
// @Accept      json
// @Produce     json
// @Param       page query     int     false   "Page number"
// @Param       limit query    int     false   "Number of countries per page"
// @Success     200  {object}  dto.CountriesWithPagination
// @Success     404  {object}  string
// @Router      /countries [get]
func (h *Handler) GetAllCountries(w http.ResponseWriter, r *http.Request) {
	pageNumber, limit := parsePaginationParams(r)

	countries, err := h.service.GetAllCountries(pageNumber, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	if len(countries) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	response := dto.CountriesWithPagination{
		Countries: countries,
		Page:      pageNumber,
		Limit:     limit,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func parsePaginationParams(r *http.Request) (pageNumber int, limit int) {
	page := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	pageNumber, _ = strconv.Atoi(page)
	if pageNumber <= 0 {
		pageNumber = 1
	}

	limit, _ = strconv.Atoi(limitParam)
	if limit <= 0 {
		limit = 10
	}

	return pageNumber, limit
}
