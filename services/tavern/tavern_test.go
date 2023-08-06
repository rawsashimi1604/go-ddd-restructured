package tavern

import (
	// "context"
	"testing"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/tavern/domain/product"
	"github.com/rawsashimi1604/tavern/services/order"
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

func Test_Tavern(t *testing.T) {
	products := initProducts(t)

	os, err := order.NewOrderService(
		// WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	customerId, err := os.AddCustomer("Gavin")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(customerId, order)

	if err != nil {
		t.Fatal(err)
	}

}
