package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/pkg/entity"
	"myapp/pkg/logger"
	"myapp/pkg/service"
	"net/http"
)

// type CountryHandlerInterface interface {
// 	GetCountries(w http.ResponseWriter, r *http.Request)
// 	EditCountry(w http.ResponseWriter, r *http.Request)
// 	PostCountry(w http.ResponseWriter, r *http.Request)
// }

// type CountryHandler struct {
// 	CountryService *service.CountryService
// }

// func NewCountryHandler(countryService *service.CountryService) *CountryHandler {
// 	return &CountryHandler{
// 		CountryService: countryService,
// 	}
// }

func PostCountry(countryService service.CountryService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		logger.LogDebug(ctx, "Entering Handler.PostCountry() ...")
		var country entity.Country
		if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
			return
		}
		countryfound, err := countryService.PostCountry(ctx, country)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error posting country: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(countryfound)
	})
}

// ///
func EditCountry(countryService service.CountryService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		logger.LogDebug(ctx, "Entering Handler.EditCountry() ...")
		var country entity.Country
		if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
			return
		}
		countryfound, err := countryService.EditCountry(ctx, country)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error editing country: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(countryfound)
	})
}

func GetCountries(countryService service.CountryService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.LogDebug(ctx, "Entering Handler.GetCountries() ...")
		pagination := entity.Pagination{
			RowsNumber: r.URL.Query().Get("size"),
			PageNumber: r.URL.Query().Get("page"),
		}
		countries, err := countryService.GetCountries(ctx, pagination)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting countries: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(countries)
	})
}

// func (h *CountryHandler) GetCountries(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	logger.LogDebug(ctx, "Entering Handler.GetCountries() ...")
// 	pagination := entity.Pagination{
// 		RowsNumber: r.URL.Query().Get("size"),
// 		PageNumber: r.URL.Query().Get("page"),
// 	}
// 	countries, err := (*h.CountryService).GetCountries(ctx, pagination)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error getting countries: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(countries)
// }
// func (h *CountryHandler) EditCountry(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	logger.LogDebug(ctx, "Entering Handler.EditCountry() ...")
// 	var country entity.Country
// 	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
// 		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
// 		return
// 	}
// 	countryfound, err := (*h.CountryService).EditCountry(ctx, country)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error editing country: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(countryfound)
// }
// func (h *CountryHandler) PostCountry(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	logger.LogDebug(ctx, "Entering Handler.PostCountry() ...")
// 	var country entity.Country
// 	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
// 		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
// 		return
// 	}
// 	countryfound, err := (*h.CountryService).PostCountry(ctx, country)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error posting country: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(countryfound)
// }
