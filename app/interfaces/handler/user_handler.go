package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/presenter"
	xvalidator "github.com/ubaidillahhf/go-clarch/app/infra/validator"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

type IUserHandler interface {
	Register(c *fiber.Ctx) error
}

type userHandler struct {
	userUsecase usecases.IUserUsecase
}

func NewUserHandler(userUsecase *usecases.IUserUsecase) IUserHandler {
	return &userHandler{
		userUsecase: *userUsecase,
	}
}

func (co *userHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := new(domain.RegisterRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	newData := domain.User{
		Email:          request.Email,
		Password:       request.Password,
		FavoritePhrase: request.FavoritePhrase,
	}
	res, resErr := co.userUsecase.Register(ctx, newData)
	if resErr != nil {
		return c.JSON(presenter.Error(resErr.Err.Error(), nil, resErr.Code))
	}

	newRes := domain.RegisterResponse{
		Id:    res.Id,
		Email: res.Email,
	}
	return c.JSON(presenter.Success("Success", newRes, nil))
}
