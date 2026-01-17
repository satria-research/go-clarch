package domain

import "context"

type ProductRepository interface {
	Insert(ctx context.Context, product Product) (Product, error)
	FindAll(ctx context.Context) ([]Product, error)
	DeleteAll(ctx context.Context) error
}

type UserRepository interface {
	Insert(ctx context.Context, newData User) (User, error)
	FindByIdentifier(ctx context.Context, username, email string) (User, error)
}
