package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"myapp/api/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	// Open DB connection
	db, err := sql.Open("mysql", "root:Berat_2005@tcp(mysql:3306)/sakila")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	// Retry pinging DB until ready
	const maxAttempts = 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		err := db.PingContext(ctx)
		if err != nil {
			log.Printf("DB not ready (attempt %d/%d): %v", attempts, maxAttempts, err)
			time.Sleep(3 * time.Second)
			continue
		}
		log.Println("Successfully connected to DB")
		break
	}

	// Final check if still not connected
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to DB after %d attempts: %v", maxAttempts, err)
	}

	// Initialize Repository, Service, Handler layers
	_, countryService := controller.InitController(ctx, db)
	// countryHandler := handler.NewCountryHandler(&countryService)
	// Setup Router
	r := mux.NewRouter()

	controller.CountryHandlers(r, countryService)
	// r.HandleFunc("/country", countryHandler.GetCountries).Methods("GET")
	// r.HandleFunc("/country", countryHandler.EditCountry).Methods("PUT")
	// r.HandleFunc("/country", countryHandler.PostCountry).Methods("POST")

	// Start HTTP Server
	port := ":8080"
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
