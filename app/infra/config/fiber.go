package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
