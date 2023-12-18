package repository

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	Insert(newData domain.User) (domain.User, *exception.Error)
}

func NewUserRepository(context context.Context, database *mongo.Database) IUserRepository {
	return &userRepository{
		Collection: database.Collection("users"),
		Context:    context,
	}
}

type userRepository struct {
	Collection *mongo.Collection
	Context    context.Context
}

func (repo *userRepository) Insert(newData domain.User) (res domain.User, err *exception.Error) {

	oneDoc := domain.User{
		Id:             gonanoid.Must(),
		Email:          newData.Email,
		Password:       newData.Password,
		FavoritePhrase: newData.FavoritePhrase,
	}

	data, dataErr := repo.Collection.InsertOne(repo.Context, oneDoc)
	if dataErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  dataErr,
		}
	}

	newId := data.InsertedID
	newData.Id = newId.(string)

	return newData, nil
}
