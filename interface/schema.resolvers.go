package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	domainModel "github.com/cpoulsen/countries/domain/model"
	"github.com/cpoulsen/countries/graph/generated"
	"github.com/cpoulsen/countries/graph/model"
)

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	var countriesSlice []*model.Country

	var dbCountries []domainModel.Country
	dbCountries = r.CountryService.GetAll()

	for _, country := range dbCountries {
		countriesSlice = append(countriesSlice, &model.Country{ID: country.ID, Name: country.Name})
	}
	return countriesSlice, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
