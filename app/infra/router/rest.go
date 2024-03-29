package router

import (
	"fmt"

	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/interfaces/handler"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

func Init(useCase usecases.AppUseCase, conf config.IConfig) {
	router := fiber.New()

	router.Get("/swagger/*", fiberSwagger.WrapHandler)

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
	router.Use(fibersentry.New(fibersentry.Config{
		Repanic: true,
	}))

	// hadler
	productHadler := handler.NewProductHandler(&useCase.ProductUsecase)
	userHandler := handler.NewUserHandler(&useCase.UserUsecase)

	// service route
	router.Get("/", handler.GetTopRoute)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/users")
	user.Post("/register", userHandler.Register)
	user.Post("/login", userHandler.Login)

	product := v1.Group("/products")
	product.Get("/", productHadler.List)
	product.Post("/", productHadler.Create)

	router.Get("/error", func(c *fiber.Ctx) error {

		type some struct {
			coba *int
		}

		newSome := new(some)

		fmt.Println(*newSome.coba)

		return c.JSON("ok")
	})

	router.Listen(":" + conf.Get("PORT"))
}
