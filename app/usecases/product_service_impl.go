package usecases

import (
	"github.com/ubaidillahhf/go-clarch/app/domain/entity"
	"github.com/ubaidillahhf/go-clarch/app/infra/model"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
	"github.com/ubaidillahhf/go-clarch/app/interfaces/validation"
)

func NewProductUsecase(productRepository *repository.IProductRepository) IProductUsecase {
	return &productUsecaseImpl{
		ProductRepository: *productRepository,
	}
}

type productUsecaseImpl struct {
	ProductRepository repository.IProductRepository
}

func (service *productUsecaseImpl) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.Validate(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.ProductRepository.Insert(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}

func (service *productUsecaseImpl) List() (responses []model.GetProductResponse) {
	products := service.ProductRepository.FindAll()
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses
}
