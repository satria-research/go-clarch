package usecases

import (
	"context"

	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"github.com/ubaidillahhf/go-clarch/app/infra/repository"
	"github.com/ubaidillahhf/go-clarch/app/infra/utility/helper"
)

type IUserUsecase interface {
	Register(ctx context.Context, request domain.User) (domain.User, *exception.Error)
}

func NewUserUsecase(repo *repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		repo: *repo,
	}
}

type userUsecase struct {
	repo repository.IUserRepository
}

func (service *userUsecase) Register(ctx context.Context, request domain.User) (res domain.User, err *exception.Error) {

	hashPwd, _ := helper.HashPassword(request.Password)
	newData := domain.User{
		Email:          request.Email,
		FavoritePhrase: request.FavoritePhrase,
		Password:       hashPwd,
	}

	p, pErr := service.repo.Insert(ctx, newData)
	if pErr != nil {
		return res, pErr
	}

	return p, nil
}
