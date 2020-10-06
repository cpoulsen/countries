package repo

import (
    "github.com/cpoulsen/countries/domain/model"
)


type Repo struct {
}

func ProvideRepo() *Repo {
  return &Repo{}
}

func (r Repo) GetAllCountries() []model.Country {
    countries := []model.Country{model.Country{Name: "Denmark", ID: "1"}, model.Country{Name: "Sweden", ID: "2"}, model.Country{Name: "Norway", ID: "3"}, model.Country{Name: "Germany", ID: "4"}, model.Country{Name: "Spain", ID: "5"}, model.Country{Name: "France", ID: "6"}}

    return countries
}
