package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/cpoulsen/countries/graph/generated"
    "github.com/cpoulsen/countries/interface"
	"github.com/cpoulsen/countries/domain/service/interface"
    "github.com/cpoulsen/countries/domain/service/implementation"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var countryService countryInterface.CountryService

	countryService = countryImplementation.NewCountryService()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
    		CountryService: countryService,
    	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
