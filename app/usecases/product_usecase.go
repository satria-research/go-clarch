package usecases

import (
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
)

type IProductUsecase interface {
	Create(request domain.CreateProductRequest) (domain.Product, *exception.Error)
	List() ([]domain.Product, *exception.Error)
}

func NewProductUsecase(productRepository *repository.IProductRepository) IProductUsecase {
	return &productUsecaseImpl{
		ProductRepository: *productRepository,
	}
}

type productUsecaseImpl struct {
	ProductRepository repository.IProductRepository
}

func (service *productUsecaseImpl) Create(request domain.CreateProductRequest) (res domain.Product, err *exception.Error) {

	newProduct := domain.Product{
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	p, pErr := service.ProductRepository.Insert(newProduct)
	if pErr != nil {
		return res, pErr
	}

	return p, nil
}

func (service *productUsecaseImpl) List() (responses []domain.Product, err *exception.Error) {
	products, pErr := service.ProductRepository.FindAll()
	if pErr != nil {
		return responses, pErr
	}

	for _, product := range products {
		responses = append(responses, domain.Product{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return responses, nil
}
