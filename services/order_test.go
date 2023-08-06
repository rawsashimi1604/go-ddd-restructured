package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/tavern/domain/customer"
	"github.com/rawsashimi1604/tavern/domain/product"
)

func initProducts(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "Some beverage", 10.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
		beer, peanuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
