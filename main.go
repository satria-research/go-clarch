package main

import (
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
	"github.com/ubaidillahhf/go-clarch/app/infra/router"
	"github.com/ubaidillahhf/go-clarch/app/infra/service"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
	_ "github.com/ubaidillahhf/go-clarch/docs"
)

// @title			Go Clarch Boilerplate
// @version			1.0
// @description		This is a golang clean architecture boilerplate using fiber
// @contact.name	Ubaidillah Hakim Fadly
// @contact.url		https://ubed.dev
// @contact.email	ubai@codespace.id
// @license.name	MIT License
// @license.url		https://github.com/satria-research/go-clarch?tab=MIT-1-ov-file
// @host			localhost
// @BasePath		/v1
//
//	@schemes                     http https
//	@securityDefinitions.apiKey  JWT
//	@in                          header
//	@name                        Authorization
//	@description                 JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {
	// Load config
	configuration := config.New(".env")

	// Error monitoring
	config.SentryInit(configuration)

	// Database connection
	database := config.NewMongoDatabase(configuration)

	// Setup repositories (Infrastructure -> Domain)
	productRepository := repository.NewProductRepository(database)
	userRepository := repository.NewUserRepository(database)

	// Setup domain services (Infrastructure -> Domain)
	passwordHasher := service.NewPasswordHasher()
	usernameGenerator := service.NewUsernameGenerator()
	tokenGenerator := service.NewTokenGenerator(configuration.Get("ACCESS_TOKEN_SECRET"))
	configProvider := service.NewConfigProvider(configuration)

	// Setup use cases (Application layer)
	useCase := usecases.NewAppUseCase(
		productRepository,
		userRepository,
		passwordHasher,
		usernameGenerator,
		tokenGenerator,
		configProvider,
	)

	// Start router (Interface layer)
	router.Init(useCase, configuration)
}
