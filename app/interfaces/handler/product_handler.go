package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/presenter"
	"github.com/ubaidillahhf/go-clarch/app/interfaces/dto"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

type ProductHandler interface {
	Create(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}

type productHandler struct {
	usecase   usecases.ProductUsecase
	validator *validator.Validate
}

func NewProductHandler(usecase usecases.ProductUsecase, validator *validator.Validate) ProductHandler {
	return &productHandler{
		usecase:   usecase,
		validator: validator,
	}
}

//		@Summary        Create new product
//		@Description    Create new product
//		@Tags           Products
//		@Accept         json
//		@Produce        json
//	 @Param data body dto.CreateProductRequest true "body payload"
//		@Security       JWT
//		@Success        200 {object} presenter.SuccessResponse{data=dto.ProductResponse}
//		@Router         /products [post]
func (h *productHandler) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var request dto.CreateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	if err := h.validator.Struct(request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	product, err := h.usecase.Create(ctx, request.Name, request.Price, request.Quantity)
	if err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.IntenalError))
	}

	response := dto.ProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return c.JSON(presenter.Success("Success", response, nil))
}

func (h *productHandler) List(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := h.usecase.List(ctx)
	if err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.IntenalError))
	}

	var responses []dto.ProductResponse
	for _, product := range products {
		responses = append(responses, dto.ProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}

	return c.JSON(presenter.Success("Success", responses, nil))
}
