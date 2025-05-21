package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"myapp/pkg/entity"
	"myapp/pkg/logger"
	"strconv"
)

// Defines the interface for the Country repository
type CountryRepository interface {
	GetCountries(ctx context.Context, pagination entity.Pagination) ([]entity.Country, error)
	EditCountry(ctx context.Context, country entity.Country) (entity.Country, error)
	PostCountry(ctx context.Context, country entity.Country) (entity.Country, error)
	GetCountryByID(ctx context.Context, id int) (entity.Country, error)
}

// MySQLCountryRepository implements the CountryRepository interface
// and provides methods to interact with the MySQL database.
type MySQLCountryRepository struct {
	db *sql.DB
}

// NewMySQLCountryRepository creates a new instance of MySQLCountryRepository
// and initializes it with the provided database connection
func NewMySQLCountryRepository(ctx context.Context, db *sql.DB) CountryRepository {
	logger.LogDebug(ctx, "Entering Repository.MySQLAssetRepository() ...")
	return &MySQLCountryRepository{db: db}
}

// GetCountryByID retrieves a country by its ID from the database
func (r *MySQLCountryRepository) GetCountryByID(ctx context.Context, id int) (entity.Country, error) {
	var country entity.Country
	query := "SELECT country_id, country, last_update FROM country WHERE country_id = ?"
	row := r.db.QueryRowContext(ctx, query, id) //Used to find a single row with that ID

	err := row.Scan(&country.ID, &country.Country, &country.LastUpdate)
	if err != nil { //
		if err == sql.ErrNoRows { // No rows were returned
			return entity.Country{}, fmt.Errorf("country not found") // Return a custom error
		}
		log.Printf("Error scanning country: %v", err)
		return entity.Country{}, err // Return the error
	}

	return country, nil // Return the found country
}

// EditCountry updates a country in the database
func (r *MySQLCountryRepository) EditCountry(ctx context.Context, country entity.Country) (entity.Country, error) {
	query := "UPDATE country SET country = ?, last_update = ? WHERE country_id = ?"
	_, err := r.db.ExecContext(ctx, query, country.Country, country.LastUpdate, country.ID)
	if err != nil {
		log.Printf("Error updating country: %v", err)
		return entity.Country{}, err
	}
	return country, nil
}

func (r *MySQLCountryRepository) PostCountry(ctx context.Context, country entity.Country) (entity.Country, error) {
	query := "INSERT INTO country (country, last_update) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx, query, country.Country, country.LastUpdate)
	if err != nil {
		log.Printf("Error inserting country: %v", err)
		return entity.Country{}, err
	}
	return country, nil
}

func (r *MySQLCountryRepository) GetCountries(ctx context.Context, pagination entity.Pagination) ([]entity.Country, error) {
	var countries []entity.Country

	pageNumber, err := strconv.Atoi(pagination.PageNumber)
	if err != nil {
		log.Printf("Error converting page number: %v", err)
		return nil, err
	}

	rowsNumber, err := strconv.Atoi(pagination.RowsNumber)
	if err != nil {
		log.Printf("Error converting rows number: %v", err)
		return nil, err
	}

	offset := (pageNumber - 1) * rowsNumber

	query := "SELECT country_id, country, last_update FROM country LIMIT ? OFFSET ?"
	rows, err := r.db.QueryContext(ctx, query, rowsNumber, offset)
	if err != nil {
		log.Printf("Error querying countries: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country entity.Country
		err := rows.Scan(&country.ID, &country.Country, &country.LastUpdate)
		if err != nil {
			log.Printf("Error scanning country: %v", err)
			return nil, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}
