package repository

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) domain.UserRepository {
	return &userRepository{
		Collection: database.Collection("users"),
	}
}

type userRepository struct {
	Collection *mongo.Collection
}

func (repo *userRepository) Insert(ctx context.Context, newData domain.User) (domain.User, error) {
	oneDoc := domain.User{
		Id:             gonanoid.Must(),
		Username:       newData.Username,
		Fullname:       newData.Fullname,
		Email:          newData.Email,
		Password:       newData.Password,
		FavoritePhrase: newData.FavoritePhrase,
	}

	data, err := repo.Collection.InsertOne(ctx, oneDoc)
	if err != nil {
		return domain.User{}, err
	}

	newId := data.InsertedID
	oneDoc.Id = newId.(string)

	return oneDoc, nil
}

func (repo *userRepository) FindByIdentifier(ctx context.Context, username, email string) (domain.User, error) {
	filter := bson.M{}
	if email != "" {
		filter["email"] = email
	}
	if username != "" {
		filter["username"] = username
	}

	var res domain.User
	err := repo.Collection.FindOne(ctx, filter).Decode(&res)
	return res, err
}
