// package aggregate holds our aggregates that combines many entities into a full object
package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/tavern"
)

var (
	ErrInvalidPerson = errors.New("A customer has to have a valid name.")
)

type Customer struct {
	// Person is the root entity of customer
	// which means person.ID is the main identifier for the customer

	// Small case -> external cannot use.
	// Not accessible directly to the other data. Not accessible from outside.
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}
