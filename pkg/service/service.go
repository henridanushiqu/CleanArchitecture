package service

import (
	"context"
	"myapp/pkg/entity"
	"myapp/pkg/logger"
	"myapp/pkg/repository"
)

type CountryService interface {
	GetCountries(ctx context.Context, pagination entity.Pagination) ([]entity.Country, error)
	EditCountry(ctx context.Context, country entity.Country) (entity.Country, error)
	PostCountry(ctx context.Context, country entity.Country) (entity.Country, error)
	GetCountryByID(ctx context.Context, id int) (entity.Country, error)
}

type MySQLCountryImpl struct {
	Repo repository.CountryRepository
}

func NewCountryService(ctx context.Context, r repository.CountryRepository) CountryService {
	return &MySQLCountryImpl{
		Repo: r,
	}
}

func (s *MySQLCountryImpl) GetCountries(ctx context.Context, pagination entity.Pagination) ([]entity.Country, error) {
	logger.LogDebug(ctx, "Entering Service.NewCountryService() ...")
	return s.Repo.GetCountries(ctx, pagination)
}

func (s *MySQLCountryImpl) EditCountry(ctx context.Context, country entity.Country) (entity.Country, error) {
	logger.LogDebug(ctx, "Entering Service.EditCountry() ...")
	countryfound, err := s.Repo.GetCountryByID(ctx, country.ID)
	if err != nil {
		return entity.Country{}, err
	}
	return s.Repo.EditCountry(ctx, countryfound)
}

func (s *MySQLCountryImpl) PostCountry(ctx context.Context, country entity.Country) (entity.Country, error) {
	logger.LogDebug(ctx, "Entering Service.PostCountry() ...")
	postCountry, err := s.Repo.GetCountryByID(ctx, country.ID)
	if err != nil {
		return entity.Country{}, err
	}
	return postCountry, err
}

func (s *MySQLCountryImpl) GetCountryByID(ctx context.Context, id int) (entity.Country, error) {
	logger.LogDebug(ctx, "Entering Service.GetCountryByID() ...")
	return s.Repo.GetCountryByID(ctx, id)
}
