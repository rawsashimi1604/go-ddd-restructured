package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/go-ddd/aggregate"
)

var (
	ErrProductNotFound      = errors.New("No such product.")
	ErrProductAlreadyExists = errors.New("Product already exists.")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
