package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/tavern/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := customer.NewCustomer("Gavin")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()
	randomId, _ := uuid.NewRandom()

	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          randomId,
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("exected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
