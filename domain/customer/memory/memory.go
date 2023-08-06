// Package memory is an in-memory implementation of CustomerRepository
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/tavern/domain/customer"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]customer.Customer, 0),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
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

func (mr *MemoryCustomerRepository) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
