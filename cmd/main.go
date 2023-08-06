package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/tavern/domain/product"
	"github.com/rawsashimi1604/tavern/services/order"
	"github.com/rawsashimi1604/tavern/services/tavern"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)

	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "Some beverage", 10.99)
	if err != nil {
		panic(err)
	}

	return []product.Product{
		beer, peanuts, wine,
	}
}
