package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/infra/model"
)

type (
	Rootmessage struct {
		ServiceName string `json:"serviceName"`
		Version     string `json:"version"`
		Author      string `json:"author"`
		Year        string `json:"year"`
	}
)

func GetTopRoute(c *fiber.Ctx) error {
	var data Rootmessage

	data.ServiceName = "GO CLARCH"
	data.Version = "1.0"
	data.Author = "Ubaidillah Hakim Fadly"
	data.Year = "2023"

	res := model.ApiResponse{
		Status: "Success",
		Data:   data,
		Code:   200,
	}
	return c.JSON(res)
}
