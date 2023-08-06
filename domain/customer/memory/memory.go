// Package memory is an in-memory implementation of CustomerRepository
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/go-ddd/aggregate"
	"github.com/rawsashimi1604/go-ddd/domain/customer"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer, 0),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure customer is already in repo
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryCustomerRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
