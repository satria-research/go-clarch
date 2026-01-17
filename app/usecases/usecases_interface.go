package usecases

import "github.com/ubaidillahhf/go-clarch/app/domain"

type AppUseCase struct {
	ProductUsecase ProductUsecase
	UserUsecase    UserUsecase
}

func NewAppUseCase(
	productRepo domain.ProductRepository,
	userRepo domain.UserRepository,
	passwordHasher domain.PasswordHasher,
	usernameGenerator domain.UsernameGenerator,
	tokenGenerator domain.TokenGenerator,
	config domain.ConfigProvider,
) AppUseCase {
	return AppUseCase{
		ProductUsecase: NewProductUsecase(productRepo),
		UserUsecase:    NewUserUsecase(userRepo, passwordHasher, usernameGenerator, tokenGenerator, config),
	}
}
