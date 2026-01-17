package usecases

import (
	"context"

	"github.com/ubaidillahhf/go-clarch/app/domain"
)

type ProductUsecase interface {
	Create(ctx context.Context, name string, price int64, quantity int32) (domain.Product, error)
	List(ctx context.Context) ([]domain.Product, error)
}

func NewProductUsecase(productRepository domain.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepository,
	}
}

type productUsecase struct {
	productRepo domain.ProductRepository
}

func (uc *productUsecase) Create(ctx context.Context, name string, price int64, quantity int32) (domain.Product, error) {
	newProduct := domain.Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	return uc.productRepo.Insert(ctx, newProduct)
}

func (uc *productUsecase) List(ctx context.Context) ([]domain.Product, error) {
	return uc.productRepo.FindAll(ctx)
}
