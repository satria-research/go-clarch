package repository

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	Insert(newData domain.User) (domain.User, *exception.Error)
}

func NewUserRepository(database *mongo.Database) IUserRepository {
	return &userRepository{
		Collection: database.Collection("users"),
	}
}

type userRepository struct {
	Collection *mongo.Collection
}

func (repository *userRepository) Insert(newData domain.User) (res domain.User, err *exception.Error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	oneDoc := domain.User{
		Id:             gonanoid.Must(),
		Email:          newData.Email,
		Password:       newData.Password,
		FavoritePhrase: newData.FavoritePhrase,
	}

	data, dataErr := repository.Collection.InsertOne(ctx, oneDoc)
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
