package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("No such product.")
	ErrProductAlreadyExists = errors.New("Product already exists.")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
