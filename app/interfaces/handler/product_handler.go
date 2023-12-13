package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/presenter"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

type ProductController struct {
	ProductService usecases.IProductUsecase
}

func NewProductController(productService *usecases.IProductUsecase) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Create(c *fiber.Ctx) error {

	var request domain.CreateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	res, resErr := controller.ProductService.Create(request)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	return c.JSON(presenter.Success("Success", res, nil))
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	responses, err := controller.ProductService.List()
	if err != nil {
		return c.JSON(presenter.Error(err.Err.Error(), nil, err.Code))
	}

	return c.JSON(presenter.Success("Success", responses, nil))
}
