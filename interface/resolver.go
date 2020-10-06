package interfaces

import (
	"github.com/cpoulsen/countries/domain/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    CountryService countryService.CountryServiceInterface
}