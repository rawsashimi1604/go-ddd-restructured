// the services package holds all services that relate to business logic
// Service Configuration Generator Pattern
package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/go-ddd/aggregate"
	"github.com/rawsashimi1604/go-ddd/domain/customer"
	"github.com/rawsashimi1604/go-ddd/domain/customer/memory"
	"github.com/rawsashimi1604/go-ddd/domain/customer/mongo"
	"github.com/rawsashimi1604/go-ddd/domain/product"
	prodmem "github.com/rawsashimi1604/go-ddd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Loop through all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each product
	var products []aggregate.Product
	var totalPrice float64

	for _, id := range productIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		totalPrice += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products. The total price is $%v.", c.GetID(), len(products), totalPrice)
	return totalPrice, nil
}
