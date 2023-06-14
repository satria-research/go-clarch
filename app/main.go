package main

import (
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
	"github.com/ubaidillahhf/go-clarch/app/infra/router"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

func main() {
	// Setup Configuration
	configuration := config.New(".env")
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(database)

	// Setup Service
	useCase := usecases.NewAppUseCase(
		productRepository,
	)

	router.Init(useCase, configuration)
}
