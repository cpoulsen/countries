package countryService

import (
    "github.com/cpoulsen/countries/domain/model"
    "github.com/cpoulsen/countries/domain/repo"
)

type CountryServiceInterface interface {
	GetAllCountries() []model.Country
}

type CountryService struct {
    repo *repo.Repo
}


func ProvideCountryService(repo *repo.Repo) *CountryService {
	return &CountryService{
	    repo,
	}
}


func (c *CountryService) GetAllCountries() []model.Country {
    return c.repo.GetAllCountries()
}