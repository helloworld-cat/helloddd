package productlistsrc

import (
	"github.com/pagedegeek/helloddd/domain/entities/product"
)

func New() *service {
	return &service{}
}

func (src *service) List() ([]*product.Entity, error) {
	// TODO: validate inputs
	// ...

	// TODO: use a product repository
	// ...

	products := []*product.Entity{
		{ID: "1", Title: "Product 1", PriceCents: 10000},
		{ID: "2", Title: "Product 2", PriceCents: 20000},
		{ID: "3", Title: "Product 3", PriceCents: 30000},
	}

	return products, nil
}

type (
	service struct {
		// repository repository
	}

	repository interface {
		// ...
	}
)
