package controller

import (
	"context"
	"database/sql"
	"myapp/api/handler"
	"myapp/pkg/repository"
	"myapp/pkg/service"

	"github.com/gorilla/mux"
)

func InitController(ctx context.Context, db *sql.DB) (repository.CountryRepository, service.CountryService) {
	countryRepo := repository.NewMySQLCountryRepository(ctx, db)
	countryService := service.NewCountryService(ctx, countryRepo)
	return countryRepo, countryService
}

func CountryHandlers(r *mux.Router, countryService service.CountryService) {
	r.Handle("/country", handler.GetCountries(countryService)).Methods("GET")
	r.Handle("/country", handler.EditCountry(countryService)).Methods("PUT")
	r.Handle("/country", handler.PostCountry(countryService)).Methods("POST")
}
