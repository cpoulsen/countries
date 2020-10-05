package interfaces

import (
	"github.com/cpoulsen/countries/domain/service/interface"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    CountryService countryInterface.CountryService
}