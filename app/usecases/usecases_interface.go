package usecases

import "github.com/ubaidillahhf/go-clarch/app/infra/repository"

type AppUseCase struct {
	ProductUsecase IProductUsecase
}

func NewAppUseCase(
	ProdukRepo repository.IProductRepository,
) AppUseCase {
	return AppUseCase{
		ProductUsecase: NewProductUsecase(&ProdukRepo),
	}
}
