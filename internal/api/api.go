package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/samuellfa/copa-do-mundo-golang/docs"
	"github.com/samuellfa/copa-do-mundo-golang/pkg/database"

	"github.com/gorilla/mux"
	"github.com/samuellfa/copa-do-mundo-golang/internal/country"
	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct{}

func New() *API {
	return &API{}
}

func (api API) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func (api API) SetupAndListen() {
	db := database.Init()

	router := mux.NewRouter()
	router.HandleFunc("/health", api.Health).Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	country.New(
		country.Config{
			Router: router,
			DB:     db,
		},
	)

	err := http.ListenAndServe(":3333", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
