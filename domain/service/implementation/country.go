package countryImplementation

import (
    "github.com/cpoulsen/countries/domain/model"
    "github.com/cpoulsen/countries/domain/repository"
)

type CountryService struct {}


func NewCountryService() *CountryService {
	return &CountryService{}
}


func (s *CountryService) GetAll() []model.Country {
    return repository.GetAll()
}