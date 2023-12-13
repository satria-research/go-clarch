package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/interfaces/handler"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

func Init(useCase usecases.AppUseCase, conf config.IConfig) {
	router := fiber.New()

	// middleware
	allowCors := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Authorization, Origin, Content-Length, Content-Type, User-Agent, Referrer, Host, Token, CSRF-Token",
		AllowMethods:     "GET, POST, PATCH, OPTIONS, PUT, DELETE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	})
	logging := logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	})
	router.Use(allowCors, logging, recover.New())

	// hadler
	productHadler := handler.NewProductController(&useCase.ProductUsecase)

	// service route
	router.Get("/", handler.GetTopRoute)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	product := v1.Group("/products")
	product.Get("/", productHadler.List)
	product.Post("/", productHadler.Create)

	router.Listen(":" + conf.Get("PORT"))
}
