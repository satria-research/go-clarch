package usecases

import "github.com/ubaidillahhf/go-clarch/app/infra/model"

type IProductUsecase interface {
	Create(request model.CreateProductRequest) (response model.CreateProductResponse)
	List() (responses []model.GetProductResponse)
}
