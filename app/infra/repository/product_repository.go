package repository

import "github.com/ubaidillahhf/go-clarch/app/domain/entity"

type IProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}
