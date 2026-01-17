package repository

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRepository(database *mongo.Database) domain.ProductRepository {
	return &productRepository{
		Collection: database.Collection("products"),
	}
}

type productRepository struct {
	Collection *mongo.Collection
}

func (repo *productRepository) Insert(ctx context.Context, product domain.Product) (domain.Product, error) {
	oneDoc := domain.Product{
		Id:       gonanoid.Must(),
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	data, err := repo.Collection.InsertOne(ctx, oneDoc)
	if err != nil {
		return domain.Product{}, err
	}

	newId := data.InsertedID
	oneDoc.Id = newId.(string)

	return oneDoc, nil
}

func (repo *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	c, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var documents []bson.M
	if err := c.All(ctx, &documents); err != nil {
		return nil, err
	}

	var res []domain.Product
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

func (repo *productRepository) DeleteAll(ctx context.Context) error {
	_, err := repo.Collection.DeleteMany(ctx, bson.M{})
	return err
}
