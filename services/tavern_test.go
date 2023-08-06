package services

import (
	// "context"
	"testing"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/go-ddd/domain/customer"
)

func Test_Tavern(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		// WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}

}
