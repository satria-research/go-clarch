package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/model"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

type ProductController struct {
	ProductService usecases.ProductUsecase
}

func NewProductController(productService *usecases.ProductUsecase) ProductController {
	return ProductController{ProductService: *productService}
}

func (handler *ProductController) Route(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.JSON(model.ApiResponse{
			Code:   200,
			Status: "OK",
			Data: map[string]string{
				"Service": "GO CLARCH IS RUNNING",
			},
		})
		return nil
	})
	app.Post("/api/products", handler.Create)
	app.Get("/api/products", handler.List)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.Create(request)
	return c.JSON(model.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	responses := controller.ProductService.List()
	return c.JSON(model.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
