package aggregate_test

import (
	"errors"
	"testing"

	"github.com/rawsashimi1604/go-ddd/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Gavin Loo",
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(testCase.name)

			if !errors.Is(err, testCase.expectedErr) {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}

}
