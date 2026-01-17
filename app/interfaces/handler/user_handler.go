package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/presenter"
	xvalidator "github.com/ubaidillahhf/go-clarch/app/infra/validator"
	"github.com/ubaidillahhf/go-clarch/app/interfaces/dto"
	"github.com/ubaidillahhf/go-clarch/app/usecases"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userHandler struct {
	usecase   usecases.UserUsecase
	validator *validator.Validate
}

func NewUserHandler(usecase usecases.UserUsecase, validator *validator.Validate) UserHandler {
	return &userHandler{
		usecase:   usecase,
		validator: validator,
	}
}

//		@Summary        Create new user
//		@Description    Create new user
//		@Tags           Users
//		@Accept         json
//		@Produce        json
//	 @Param data body dto.RegisterRequest true "body payload"
//		@Security       JWT
//		@Success        200 {object} presenter.SuccessResponse{data=dto.RegisterResponse}
//		@Router         /users/register [post]
func (h *userHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := new(dto.RegisterRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	if err := h.validator.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	user, err := h.usecase.Register(ctx, request.Username, request.Fullname, request.Email, request.Password, request.FavoritePhrase)
	if err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	response := dto.RegisterResponse{
		Id:    user.Id,
		Email: user.Email,
	}

	return c.JSON(presenter.Success("Success", response, nil))
}

//		@Summary        Login user
//		@Description    Login user
//		@Tags           Users
//		@Accept         json
//		@Produce        json
//	 @Param data body dto.LoginRequest true "body payload"
//		@Success        200 {object} presenter.SuccessResponse{data=dto.LoginResponse}
//		@Router         /users/login [post]
func (h *userHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := new(dto.LoginRequest)
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	if err := h.validator.Struct(request); err != nil {
		return c.JSON(presenter.Error("error", xvalidator.GenerateHumanizeError(request, err), exception.BadRequestError))
	}

	token, err := h.usecase.Login(ctx, request.Email, request.Password)
	if err != nil {
		return c.JSON(presenter.Error(err.Error(), nil, exception.BadRequestError))
	}

	response := dto.LoginResponse{
		Email: request.Email,
		Token: token,
	}

	return c.JSON(presenter.Success("Success", response, nil))
}
