package repository

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	Insert(product domain.Product) (domain.Product, *exception.Error)
	FindAll() ([]domain.Product, *exception.Error)
	DeleteAll() *exception.Error
}

func NewProductRepository(context context.Context, database *mongo.Database) IProductRepository {
	return &productRepository{
		Context:    context,
		Collection: database.Collection("products"),
	}
}

type productRepository struct {
	Collection *mongo.Collection
	Context    context.Context
}

func (repo *productRepository) Insert(product domain.Product) (res domain.Product, err *exception.Error) {

	oneDoc := domain.Product{
		Id:       gonanoid.Must(),
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	data, dataErr := repo.Collection.InsertOne(repo.Context, oneDoc)
	if dataErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  dataErr,
		}
	}

	newId := data.InsertedID
	product.Id = newId.(string)

	return product, nil
}

func (repo *productRepository) FindAll() (res []domain.Product, err *exception.Error) {

	c, cErr := repo.Collection.Find(repo.Context, bson.M{})
	if cErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  cErr,
		}
	}

	var documents []bson.M
	if err := c.All(repo.Context, &documents); err != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	for _, document := range documents {
		res = append(res, domain.Product{
			Id:       document["_id"].(string),
			Name:     document["name"].(string),
			Price:    document["price"].(int64),
			Quantity: document["quantity"].(int32),
		})
	}

	return res, nil
}

func (repo *productRepository) DeleteAll() *exception.Error {

	if _, err := repo.Collection.DeleteMany(repo.Context, bson.M{}); err != nil {
		return &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	return nil
}
