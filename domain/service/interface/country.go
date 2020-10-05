package countryInterface

import(
	"github.com/cpoulsen/countries/domain/model"
)

type CountryService interface {
	GetAll() []model.Country
}