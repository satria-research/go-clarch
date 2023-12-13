package repository

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	Insert(product domain.Product) (domain.Product, *exception.Error)
	FindAll() ([]domain.Product, *exception.Error)
	DeleteAll() *exception.Error
}

func NewProductRepository(database *mongo.Database) IProductRepository {
	return &productRepository{
		Collection: database.Collection("products"),
	}
}

type productRepository struct {
	Collection *mongo.Collection
}

func (repository *productRepository) Insert(product domain.Product) (res domain.Product, err *exception.Error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	oneDoc := domain.Product{
		Id:       gonanoid.Must(),
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	data, dataErr := repository.Collection.InsertOne(ctx, oneDoc)
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

func (repository *productRepository) FindAll() (res []domain.Product, err *exception.Error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	c, cErr := repository.Collection.Find(ctx, bson.M{})
	if cErr != nil {
		return res, &exception.Error{
			Code: exception.IntenalError,
			Err:  cErr,
		}
	}

	var documents []bson.M
	if err := c.All(ctx, &documents); err != nil {
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

func (repository *productRepository) DeleteAll() *exception.Error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	if _, err := repository.Collection.DeleteMany(ctx, bson.M{}); err != nil {
		return &exception.Error{
			Code: exception.IntenalError,
			Err:  err,
		}
	}

	return nil
}
