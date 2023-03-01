package country

import (
	"github.com/gorilla/mux"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/handler"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/repository"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country/service"
	"gorm.io/gorm"
)

type API struct {
	Handler *handler.Handler
}

type Config struct {
	Router *mux.Router
	DB     *gorm.DB
}

func New(c Config) *API {
	repository := repository.New(c.DB)
	service := service.New(repository)

	handler := handler.New(
		service,
	)

	setRoutes(
		handler,
		c.Router,
	)

	return &API{
		Handler: handler,
	}
}

func setRoutes(handler *handler.Handler, router *mux.Router) {
	r := router.PathPrefix("/v1").Subrouter()

	r.HandleFunc("/country", handler.CreateCountry).Methods("POST")
	r.HandleFunc("/country/{id}", handler.GetCountryById).Methods("GET")
	r.HandleFunc("/country", handler.GetCountryByName).Methods("GET")
}
